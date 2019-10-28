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

// categoryid 排序重载
type categoryid []*rconst.Category

func (a categoryid) Len() int {
	return len(a)
}

func (a categoryid) Less(i, j int) bool {
	if a[i].CategoryID < a[j].CategoryID {
		return true
	}
	return false
}

func (a categoryid) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
