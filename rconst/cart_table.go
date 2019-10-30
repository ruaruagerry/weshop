package rconst

// CartSpecification 货物规格
type CartSpecification struct {
	SpecificationID string `json:"specificationid"`
	DetailID        string `json:"detailid"`
	Value           string `json:"value"`
}

// Cart 购物车货物
type Cart struct {
	CartID        string               `json:"cartid"`
	GoodID        string               `json:"goodid"`
	Num           int32                `json:"num"`
	Specification []*CartSpecification `json:"specification"`
}

const (
	// ZSetCartInfoPrefix 玩家购物车+playerid
	ZSetCartInfoPrefix = "weshop:cart:info:"
)
