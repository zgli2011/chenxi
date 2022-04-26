package db

import (
	"chenxi/utils"
)

type BaseModel struct {
	ID        uint         `gorm:"column:id;primary_key,AUTO_INCREMENT;" json:"id"`                   // 主键id
	Creator   string       `gorm:"column:creator;type:varchar(64);" json:"creator"`                   // 创建人
	Updater   string       `gorm:"column:updater;type:varchar(64);" json:"updater"`                   // 更新人
	CreatedAt utils.MyTime `gorm:"column:create_time;type:datetime;default:null;" json:"create_time"` // 创建时间
	UpdatedAt utils.MyTime `gorm:"column:update_time;type:datetime;default:null;" json:"update_time"` // 变更时间
	DeletedAt utils.MyTime `gorm:"column:delete_time;type:datetime;default:null;" json:"delete_time"` // 删除时间
}

type CMDBCloudVendor struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64);" json:"name"`
	Remark string `gorm:"column:remark;type:varchar(128);" json:"remark"`
}

type CMDBCloudAccount struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(64);" json:"name"`
	Key    string `gorm:"column:key;type:varchar(256);" json:"key"`
	Secret string `gorm:"column:secret;type:varchar(256);" json:"secret"`
	Remark string `gorm:"column:remark;type:varchar(128);" json:"remark"`
}

type CMDBRegion struct {
	BaseModel
	RegionID        string       `gorm:"column:region_id;type:varchar(64);" json:"region_id"`
	RegionName      string       `gorm:"column:region_name;type:varchar(64);" json:"region_name"`
	Endpoint        string       `gorm:"column:endpoint;type:varchar(256);" json:"endpoint"`
	CheckTime       utils.MyTime `gorm:"column:check_time;type:datetime;default:null;" json:"check_time"`
	DataHash        string       `gorm:"column:data_hash;type:varchar(256);" json:"data_hash"`
	CMDBCloudVendor CMDBCloudVendor
}

type CMDBAvailableZone struct {
	BaseModel
	ZoneID           string       `gorm:"column:zone_id;type:varchar(64);" json:"zone_id"`
	ZoneName         string       `gorm:"column:zone_name;type:varchar(64);" json:"zone_name"`
	RegionID         string       `gorm:"column:region_id;type:varchar(64);" json:"region_id"`
	CheckTime        utils.MyTime `gorm:"column:check_time;type:datetime;default:null;" json:"check_time"`
	DataHash         string       `gorm:"column:data_hash;type:varchar(256);" json:"data_hash"`
	CMDBRegion       CMDBRegion
	CMDBCloudAccount CMDBCloudAccount
}

type CMDBVPC struct {
	BaseModel
}

type CMDBSubnet struct {
	BaseModel
}

type CMDBSecurityGroup struct {
	BaseModel
}

type CMDBImage struct {
	BaseModel
}

type CMDBInstanceType struct {
	BaseModel
}

type CMDBInstance struct {
	BaseModel
}

type CMDBDisk struct {
	BaseModel
}

type CMDBSnapshot struct {
	BaseModel
}

type CMDBPublicIP struct {
	BaseModel
}

type CMDBOSS struct {
	BaseModel
}

type CMDBKubernatesCluster struct {
	BaseModel
}

type CMDBKubernatesNodeGroup struct {
	BaseModel
}
