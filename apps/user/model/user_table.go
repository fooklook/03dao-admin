package model

type UserTable struct {
	PkID              uint64 `gorm:"column:pk_id;primary_key;type:bigint(20) unsigned not null auto_increment"`
	UserName          string `gorm:"column:user_name;type:varchar(64) not null"`
	NickName          string `gorm:"column:nick_name;type:varchar(64) not null"`
	WechatAccessToken string `gorm:"column:wechat_token;type:varchar(64)"`      //网页授权接口调用凭证,注意
	WechatExpiresIn   string `gorm:"column:wechat_expires_in;type:varchar(64)"` // access_token接口调用凭证超时时间，单位（秒）
	WechatOpenid      string `gorm:"column:wechat_open_id;type:varchar(64)"`    // 用户唯一标识
}
