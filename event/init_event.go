package event

import (

	"TigerScan/common"
)

const (
	SettingPath string = "conf/setting.json"
)

type AssetsAddrSetting struct {
	Addr	string	`json:"assets_db_addr"`
	Port	string	`json:"assets_db_port"`
	User	string	`json:"assets_db_user"`
	Pass	string	`json:"assets_db_pass"`
	Name	string	`json:"assets_db_name"`
	Collect	string	`json:"assets_db_collect"`

}



type SettingConf struct {
	OnlineDBAddr	OnlineDBStruct		`json:"online_db"`
	OnlineAssetsDB  AssetsAddrSetting       `json:"assets_db_info"`
	LocalPwdPath	string			`json:"dict_passwd_path"`
	LocalUserPath	string			`json:"dict_user_path"`
}

type OnlineDBStruct struct {
	DictPwdDBAddr    string `json:"passwd_dict_db_addr"`
	DictPwdDBName    string `json:"passwd_dict_db_name"`
	DictPwdDBUser    string `json:"passwd_dict_db_user"`
	DictPwdDBPass    string `json:"passwd_dict_db_pass"`
	DictPwdDBColName string `json:"passwd_dict_collect_name"`
}

var bdssDs common.DataStore

func GetDs() common.DataStore {
	return bdssDs
}

//InitPath : 初始化Mongodb 连接环境
func InitMongoAddrPath(dbName, mgoURL string) {
	ds := common.DataStore{
		DB: dbName,
		//MgoUrl:  "mongodb://bdss:1q2w3e4r@localhost/bdss",
		MgoUrl:  mgoURL,
		Session: nil,
	}
	ds.MgoInit()
	bdssDs = ds
}
