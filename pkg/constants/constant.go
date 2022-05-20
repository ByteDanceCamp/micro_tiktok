package constants

const (
	TablePre        = "ds_"
	UserServiceName = "user_service"
	UserTableName   = TablePre + "user"
	UserSalt        = "ByteDanceCamp"
	EtcdAddress     = "127.0.0.1:2379"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:13306)/gorm?charset=utf8&parseTime=True&loc=Local"
)
