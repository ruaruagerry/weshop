package rconst

// Good 商品
type Good struct {
	GoodID   string `json:"goodid"`
	URL      string `json:"url"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Price    int64  `json:"price"`
}

const (
	// HashGoodsInfoPrefix 商品列表
	HashGoodsInfoPrefix = "weshop:goods:info:"
)
