// Package rconst 定义所有用到的redis keys/fields
package rconst

const (
	// HashAccountPrefix account hash table
	HashAccountPrefix = "gd:acc:"
	// FieldAccUserID user id
	FieldAccUserID = "id"
	// FieldAccID acc id
	FieldAccID = "accId"
	// FieldAccName account name
	FieldAccName = "nickName"
	// FieldAccGender 性别
	FieldAccGender = "gender"
	// FieldAccImage 头像
	FieldAccImage = "avatarURL"
	// FieldAccOpenID account openID
	FieldAccOpenID = "openID"
	// FieldAccOs account 操作系统
	FieldAccOs = "os"
	// FiledAccCreateTime 创建时间
	FiledAccCreateTime = "createtime"
	// FiledAccLoginTime 最后登录时间
	FiledAccLoginTime = "logintime"
	// FieldAccPhoneNum 绑定手机
	FieldAccPhoneNum = "phoneNum"
)
