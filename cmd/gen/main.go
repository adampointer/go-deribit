package main

import (
	"github.com/iancoleman/strcase"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const apiUrl = "https://test.deribit.com/api/v2"

type data struct {
	Name, Method, FuncName, Args, Format, Type, Channel string
}

var subscriptions = map[string]struct {
	funcName, notificationType string
}{
	"book.{instrument_name}.{group}.{depth}.{interval}": {"BookGroup", "BookInstrumentNameGroupDepthIntervalRepeated"},
	"book.{instrument_name}.{interval}":                 {"BookInterval", "BookInstrumentIntervalRepeated"},
	"deribit_price_index.{index_name}":                  {"DeribitPriceIndex", "DeribitPriceIndexIndexNameRepeated"},
	"deribit_price_ranking.{index_name}":                {"DeribitPriceRanking", "DeribitPriceRankingIndexNameRepeated"},
	"estimated_expiration_price.{index_name}":           {"EstimatedExpirationPrice", "EstimatedExpirationPriceIndexNameRepeated"},
	"markprice.options.{index_name}":                    {"MarkPriceOptions", "MarkpriceOptionsIndexNameRepeated"},
	"perpetual.{instrument_name}.{interval}":            {"Perpetual", "PerpetualInstrumentNameIntervalRepeated"},
	"quote.{instrument_name}":                           {"Quote", "QuoteInstrumentNameRepeated"},
	"ticker.{instrument_name}.{interval}":               {"Ticker", "TickerInstrumentNameIntervalRepeated"},
	"trades.{instrument_name}.{interval}":               {"Trades", "TradesInstrumentNameIntervalRepeated"},
	"user.orders.{instrument_name}.{interval}":          {"UserOrdersInstrumentName", "UserOrdersInstrumentNameIntervalRepeated"},
	"user.orders.{kind}.{currency}.{interval}":          {"UserOrdersKind", "UserOrdersKindCurrencyIntervalRepeated"},
	"user.portfolio.{currency}":                         {"UserPortfolio", "UserPortfolioCurrencyRepeated"},
	"user.trades.{instrument_name}.{interval}":          {"UserTradesInstrument", "UserTradesInstrumentNameIntervalRepeated"},
	"user.trades.{kind}.{currency}.{interval}":          {"UserTradesKind", "UserTradesKindCurrencyIntervalRepeated"},
	"announcements":                                     {"Announcements", "Announcements"},
}

type Method struct {
	Name     string `json:"name" mapstructure:"name"`
	Category string `json:"category" mapstructure:"category"`
}

func main() {
	var d data
	var t *template.Template
	flag.StringVar(&d.Method, "method", "", "The method e.g. private/cancel_all_by_currency.")
	flag.StringVar(&d.Type, "type", "", "The response type e.g. #/definitions/private_cancel_response")
	subs := flag.Bool("subscriptions", false, "Write out subscriptions code")
	flag.Parse()
	if *subs {
		for c, params := range subscriptions {
			d.Channel = c
			d.FuncName = "Subscribe" + params.funcName
			d.Type = params.notificationType
			re := regexp.MustCompile(`\{(.*?)\}`)
			match := re.FindAllStringSubmatch(c, -1)
			if len(match) > 0 {
				args := make([]string, len(match))
				for i, m := range match {
					args[i] = m[1]
				}
				d.Args = strings.Join(args, ", ")
				d.Format = re.ReplaceAllString(c, "%s")
			} else {
				d.Format = c
			}
			t = template.Must(template.New("subMethod").Parse(subscriptionTemplate))
			err := t.Execute(os.Stdout, d)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		d.Name = strcase.ToCamel(strings.Replace(d.Method, "/", "_", -1))
		parts := strings.Split(d.Type, "/")
		d.Type = strcase.ToCamel(parts[2])
		t = template.Must(template.New("rpcMethod").Parse(methodTemplate))
		err := t.Execute(os.Stdout, d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var methodTemplate = `

// {{.Name}} makes a request to {{.Method}}
func (e *Exchange) {{.Name}}(params *operations.Get{{.Name}}Parameters) (*models.{{.Type}}, error) {
	res, err := e.rpcRequest("{{.Method}}", params)
	if err != nil {
		return nil, err
	}
	var ret models.{{.Type}}
	var toDecode interface{}
	rt := reflect.TypeOf(ret)
	switch rt.Kind() {
	case reflect.Struct:
		toDecode = res.Result
	case reflect.Slice:
		slc := make([]interface{}, len(res.Result))
		for i, j := range res.Result {
			n, err := strconv.Atoi(i)
			if err != nil {
				return nil, fmt.Errorf("error decoding result: %s", err)
			}
			slc[n] = j
		}
		toDecode = slc
	}
	if err := mapstructure.Decode(toDecode, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}`

var subscriptionTemplate = `
// {{.FuncName}} subscribes to the {{.Channel}} channel
func (e *Exchange) {{.FuncName}}({{.Args}}{{if .Args}} string{{end}}) (chan *notifications.{{.Type}}, error) {
	chans := []string{ {{if .Args}}fmt.Sprintf("{{.Format}}", {{.Args}} ) {{else}}"{{.Format}}"{{end}} }
	if _, err := e.PublicSubscribe(&public.SubscribeRequest{Channels: chans}); err != nil {
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}
	c := make(chan *RPCNotification)
	out := make(chan *notifications.{{.Type}})
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				var ret notifications.{{.Type}}
				if err := mapstructure.Decode(n.Params.Data, &ret); err != nil {
					e.errors <- fmt.Errorf("error decoding notification: %s", err)
				}
				out <- &ret
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}
`
