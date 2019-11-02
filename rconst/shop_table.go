package rconst

// ShopAddress 收货地址
type ShopAddress struct {
	AddressID string `json:"addressid"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
}

const (
	// HashShopCheckoutConfig 商城下单配置
	HashShopCheckoutConfig = "weshop:shop:checkoutconfig"
	// FIeldShopFreightPrice 运费
	FIeldShopFreightPrice = "freightprice"

	// HashShopStoreAddressPrefix 商城已保存的下单地址+playerid
	HashShopStoreAddressPrefix = "weshop:shop:storeaddress:"
	// HashShopUsedAddress 商城选择使用的下单地址
	HashShopUsedAddress = "weshop:shop:usedaddress"
)
