package rconst

// Nav 导航条
type Nav struct {
	ID        int32  `json:"id"`        // 分类ID
	Name      string `json:"name"`      // 分类名称
	BannerURL string `json:"bannerurl"` // 长条图标
}

// Category 分类
type Category struct {
	ID         int32  `json:"id"`         // 商品分类ID
	CategoryID string `json:"categoryid"` // 商品id
	Name       string `json:"name"`       // 商品分类名称
	Describe   string `json:"describe"`   // 商品分类描述
	URL        string `json:"url"`        // 商品图片
}

const (
	// HashCatalogNav 导航详情
	HashCatalogNav = "weshop:catalog:nav"
	// HashCatalogCategoryPrefix 分类详情+navid
	HashCatalogCategoryPrefix = "weshop:catalog:category:"
	// HashCatalogNavIndex 用户当前索引
	HashCatalogNavIndex = "weshop:catalog:navindex"
)
