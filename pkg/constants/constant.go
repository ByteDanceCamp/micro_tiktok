package constants

const (
	TablePre            = "ds_"
	APIServiceName      = "api_service"
	UserServiceName     = "user_service"
	RelationServiceName = "relation_service"
	CommentServiceName  = "comment_service"

	UserTableName          = TablePre + "user"
	RelationCountTableName = TablePre + "relation_count"
	FollowTableName        = TablePre + "follow"
	FollowerTableName      = TablePre + "follower"
	CommentTableName       = TablePre + "comment"
	CommentCountTableName  = TablePre + "comment_count"

	RelationFollowPre = "follow:"
	RelationFansPre   = "fans:"
	RelationCountPre  = "count:"

	UserSalt     = "ByteDanceCamp"
	JWTSecretKey = "ByteDanceCamp3"
	IdentityKey  = "uid"

	EtcdAddress     = "127.0.0.1:12379"
	RedisAddress    = "127.0.0.1:16379"
	MySQLDefaultDSN = "mtt:4&o4csZZ^OMDiy1Q@tcp(localhost:13306)/micro_tiktok?charset=utf8&parseTime=True&loc=Local"
)
