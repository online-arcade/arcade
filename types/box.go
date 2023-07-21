package types

import "time"

type Box struct {
	Id       int64     `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`     //名称
	Desc     string    `json:"desc,omitempty"`     //说明
	Icon     string    `json:"icon,omitempty"`     //图片
	Type     string    `json:"type,omitempty"`     //类型
	Disabled bool      `json:"disabled,omitempty"` //禁用
	Created  time.Time `json:"created" xorm:"created"`
}
