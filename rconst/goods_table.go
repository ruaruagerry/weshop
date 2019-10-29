package rconst

// Good 商品
type Good struct {
	GoodID   string `json:"goodid"`
	URL      string `json:"url"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Price    int64  `json:"price"`
}

// GoodAttribute 商品参数
type GoodAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GoodIssue 商品常见问题
type GoodIssue struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// GoodBrand 商品商标
type GoodBrand struct {
	BrandID   string `json:"brandid"`
	BrandName string `json:"brandname"`
}

const (
	// HashGoodsInfoPrefix 商品列表+categoryid
	HashGoodsInfoPrefix = "weshop:goods:info:"

	// HashGoodsGallery 商品展览图
	HashGoodsGallery = "weshop:goods:gallery"
	// HashGoodsAttribute 商品属性
	HashGoodsAttribute = "weshop:goods:attribute"
	// HashGoodsIssue 商品常见问题
	HashGoodsIssue = "weshop:goods:issue"
	// HashGoodsBrand 商品牌子
	HashGoodsBrand = "weshop:goods:brand"
)
