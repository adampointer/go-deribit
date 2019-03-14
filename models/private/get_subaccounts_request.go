package private

type GetSubaccountsRequest struct {
	WithPortfolio bool `json:"with_portfolio" mapstructure:"with_portfolio"`
}
