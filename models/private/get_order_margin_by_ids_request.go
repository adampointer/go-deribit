package private

type GetOrderMarginByIdsRequest struct {
	Ids []string `json:"ids" mapstructure:"ids"`
}
