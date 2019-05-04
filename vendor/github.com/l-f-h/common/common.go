package common

const (
	//1000~2000 base
	RpcSuccessful         = 1000
	RpcFailed             = 1001
	ServerInternalError   = 1002
	CallBackSuccessful    = 1003
	CallBackMarshalFailed = 1004
	CallBackPostFailed    = 1005
	//0~1000 judge
	AC int32 = 0
	EA int32 = 1
	TL int32 = 2
	ML int32 = 3
	CE int32 = 4
	RE int32 = 5
	OL int32 = 6
)

var Message = map[int32]string{
	AC:                    "通过",
	EA:                    "错误的解答",
	TL:                    "超出时间限制",
	ML:                    "内存超出限制",
	OL:                    "输出超出限制",
	CE:                    "编译错误",
	RE:                    "运行时错误(div zero/segment fault...)",
	RpcFailed:             "Rpc Failed",
	ServerInternalError:   "ServerInternalError",
	CallBackMarshalFailed: "回调时json序列化失败",
	CallBackSuccessful:    "回调成功",
	CallBackPostFailed:    "回调时发起Post失败",
}
