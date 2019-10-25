package rconst

// Nav 导航条
type Nav struct {
	ID        int32  `json:"id"`        // 分类ID
	Name      string `json:"name"`      // 分类名称
	BannerURL string `json:"bannerurl"` // 长条图标
}

// Catagory 商品
type Catagory struct {
	ID   int32  `json:"id"`   // 商品分类ID
	Name string `json:"name"` // 商品名称
	URL  string `json:"url"`  // 商品图片
}

const (
	// HashCatagoryNav 分类详情
	HashCatagoryNav = "weshop:catagory:nav"
	// HashCatagoryInfoPrefix 商品详情+navid
	HashCatagoryInfoPrefix = "weshop:catagory:info:"
	// HashCatagoryNavIndex 用户当前索引
	HashCatagoryNavIndex = "weshop:catagory:navindex"
)
