package rconst

// ShopAddress 收货地址
type ShopAddress struct {
	AddressID string `json:"addressid"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isdefault"`
	Time      int64  `json:"time"`
}

// ShopOrder 订单信息
type ShopOrder struct {
	OrderID    string  `json:"orderid"`
	AddressID  string  `json:"addressid"`
	CartIndexs []int32 `json:"cartindexs"`
}

const (
	// HashShopCheckoutConfig 商城下单配置
	HashShopCheckoutConfig = "weshop:shop:checkoutconfig"
	// FieldShopFreightPrice 运费
	FieldShopFreightPrice = "freightprice"
	// FieldShopWXPayURL 微信收钱码
	FieldShopWXPayURL = "wxpayurl"
	// FieldShopZFBPayURL 微信收钱码
	FieldShopZFBPayURL = "zfbpayurl"

	// HashShopStoreAddressPrefix 商城已保存的下单地址+playerid
	HashShopStoreAddressPrefix = "weshop:shop:storeaddress:"
	// HashShopUsedAddress 商城选择使用的下单地址
	HashShopUsedAddress = "weshop:shop:usedaddress"

	// HashShopOrderPrefix 订单地址+playerid
	HashShopOrderPrefix = "weshop:shop:order:"
)
