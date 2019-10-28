package goods

import "weshop/rconst"

// goodid 排序重载
type goodid []*rconst.Good

func (a goodid) Len() int {
	return len(a)
}

func (a goodid) Less(i, j int) bool {
	if a[i].GoodID < a[j].GoodID {
		return true
	}
	return false
}

func (a goodid) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
