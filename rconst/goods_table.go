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

// GoodSpecificationDetail 商品规格详情
type GoodSpecificationDetail struct {
	DetailID string `json:"detailid"`
	Value    string `json:"value"`
}

// GoodSpecification 商品规格
type GoodSpecification struct {
	SpecificationID string                     `json:"specificationid"`
	Name            string                     `json:"name"`
	ValueList       []*GoodSpecificationDetail `json:"valuelist"`
}

const (
	// SetGoodsCategoryPrefix 商品列表+categoryid
	SetGoodsCategoryPrefix = "weshop:goods:category:"
	// HashGoodsInfo 商品信息
	HashGoodsInfo = "wehshop:goods:info"
	// HashGoodsGallery 商品展览图
	HashGoodsGallery = "weshop:goods:gallery"
	// HashGoodsAttribute 商品属性
	HashGoodsAttribute = "weshop:goods:attribute"
	// HashGoodsIssue 商品常见问题
	HashGoodsIssue = "weshop:goods:issue"
	// HashGoodsBrand 商品牌子
	HashGoodsBrand = "weshop:goods:brand"
	// HashGoodsSpecification 商品规格
	HashGoodsSpecification = "weshop:goods:specification"
)
