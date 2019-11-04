package shop

import "weshop/rconst"

// sortaddressid 排序重载
type sortaddressid []*rconst.ShopAddress

func (a sortaddressid) Len() int {
	return len(a)
}

func (a sortaddressid) Less(i, j int) bool {
	if a[i].Time < a[j].Time {
		return true
	}
	return false
}

func (a sortaddressid) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
