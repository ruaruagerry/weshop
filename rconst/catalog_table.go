package rconst

// Nav 导航条
type Nav struct {
	ID        int32  `json:"id"`        // 分类ID
	Name      string `json:"name"`      // 分类名称
	BannerURL string `json:"bannerurl"` // 长条图标
}

// Catagory 商品
type Catagory struct {
	ID         int32  `json:"id"`         // 商品分类ID
	CatagoryID string `json:"catagoryid"` // 商品id
	Name       string `json:"name"`       // 商品名称
	URL        string `json:"url"`        // 商品图片
}

const (
	// HashCatalogNav 分类详情
	HashCatalogNav = "weshop:catalog:nav"
	// HashCatalogInfoPrefix 商品详情+navid
	HashCatalogInfoPrefix = "weshop:catalog:info:"
	// HashCatalogNavIndex 用户当前索引
	HashCatalogNavIndex = "weshop:catalog:navindex"
)
