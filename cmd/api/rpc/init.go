package rpc

// InitRPC init the RPC client
func InitRPC() {
	initUserRpc()
	initRelationRPC()
	initFavoriteRPC()
}
