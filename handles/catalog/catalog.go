package catalog

import (
	"weshop/rconst"
)

// navid 排序重载
type navid []*rconst.Nav

func (a navid) Len() int {
	return len(a)
}

func (a navid) Less(i, j int) bool {
	if a[i].ID < a[j].ID {
		return true
	}
	return false
}

func (a navid) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// catagoryid 排序重载
type catagoryid []*rconst.Catagory

func (a catagoryid) Len() int {
	return len(a)
}

func (a catagoryid) Less(i, j int) bool {
	if a[i].ID < a[j].ID {
		return true
	}
	return false
}

func (a catagoryid) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
