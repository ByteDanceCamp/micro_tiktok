package constants

const (
	TablePre            = "ds_"
	APIServiceName      = "api_service"
	UserServiceName     = "user_service"
	VideoServiceName    = "video_service"
	RelationServiceName = "relation_service"
	FavoriteServiceName = "favorite_service"
	UserTableName       = TablePre + "user"
	UserSalt            = "ByteDanceCamp"
	JWTSecretKey        = "ByteDanceCamp3"
	IdentityKey         = "uid"
	EtcdAddress         = "127.0.0.1:12379"
	RedisAddress        = "127.0.0.1:16379"
	MySQLDefaultDSN     = "mtt:4&o4csZZ^OMDiy1Q@tcp(localhost:13306)/micro_tiktok?charset=utf8&parseTime=True&loc=Local"
	RelationFollowPre   = "follow:"
	RelationFansPre     = "fans:"
	FavoriteLikePre     = "like:"
	FavoriteVideoPre    = "video:"
)
