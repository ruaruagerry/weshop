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

func statusToString(status int32) string {
	switch status {
	case rconst.TypeOrderStatusNotCharged:
		return "未付款"
	case rconst.TypeOrderStatusCharged:
		return "已付款"
	case rconst.TypeOrderStatusSend:
		return "已发货"
	case rconst.TypeOrderStatusFinished:
		return "已完成"
	}

	return "未知状态"
}
