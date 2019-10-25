package tables

/*
 * @Author: zheng
 * @Date: 2019-08-09 09:51:22
 * @Description:
 */

import (
	"encoding/json"
	"time"
)

var (
	// ChannelTest 测试
	ChannelTest = "test"
	// ChannelWeiXin 微信
	ChannelWeiXin = "104"
	// ChannelWanBa 玩吧
	ChannelWanBa = "105"
	// ChannelWeiXinPublic 微信公众号渠道
	ChannelWeiXinPublic = "106"
	// ChannelTouTiao 今日头条渠道
	ChannelTouTiao = "107"
	// ChannelYiJie 安卓包渠道 (易接平台)
	ChannelYiJie = "108"
	// ChannelIOSApp ios渠道
	ChannelIOSApp = "109"
	// ChannelAndroidApp 安卓包渠道
	ChannelAndroidApp = "110"
	// ChannelBiliBili bilibili渠道
	ChannelBiliBili = "111"
	// ChannelRobot 机器人
	ChannelRobot = "robot"
	// PFAndroid 安卓平台
	PFAndroid = "android"
	// PFIOS ios平台
	PFIOS = "ios"
)

// Account 玩家
type Account struct {
	ID            int64     `xorm:"id"  json:"id"`                      // 用户ID
	AccountID     int64     `xorm:"account_id" json:"accountId"`        // 账户ID
	GameID        int32     `xorm:"game_id" json:"gameId"`              // 游戏ID
	Nick          string    `xorm:"nick" json:"nick"`                   // 昵称
	Gender        int32     `xorm:"gender" json:"sex"`                  // 性别(注:账号服返回的性别字段为sex)
	Portrait      string    `xorm:"portrait" json:"portrait"`           // 头像
	Os            string    `xorm:"os" json:"OS"`                       // 用户操作系统
	Phoneno       string    `json:"Phoneno"`                            // 手机号码
	OpenID        string    `xorm:"open_id" json:"openId"`              // OpenID
	PublicOpenID  string    `xorm:"public_open_id" json:"publicOpenId"` // 公众号OpenId
	UnionID       string    `xorm:"union_id" json:"unionId"`            // 微信UnionID
	SessionKey    string    `xorm:"session_key" json:"sessionKey"`      // 微信SessionKey
	CreateTime    time.Time `xorm:"created" json:"CreateTime"`          // 创建时间
	LastLoginTime time.Time `xorm:"updated" json:"LastLoginTime"`       // 最后登录时间
}

// Encode 编码
func (a *Account) Encode() string {
	buf, err := json.Marshal(a)
	if err != nil {
		return "{}"
	}
	return string(buf)
}

// Decode 解码
func (a *Account) Decode(data string) error {
	err := json.Unmarshal([]byte(data), a)
	if err != nil {
		return err
	}
	return nil
}
