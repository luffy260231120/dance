package cons

const (
	ERR_PUB_UNKOWN = 200000 //未知错误
	ERR_PUB_PARAMS = 200001 //参数错误
	ERR_PUB_BUSY   = 200002 //系统繁忙
	ERR_PUB_AUTH   = 200003 //鉴权失败
	ERR_PUB_UNREG  = 200004 //设备未激活
	ERR_PUB_UNIMPL = 200005 //接口未支持
	ERR_PUB_SYSTEM = 200006 //系统错误
	ERR_PUB_RIGHT  = 200007 //无权调用接口
	ERR_NO_DATA    = 200008 //暂无数据（上游无数据）
)

type Err struct {
	Code string
	Msg  string
}

const (
	SUCCESS              = "SUCCESS"
	ERR_SIGN             = "ERR_SIGN"
	ERR_REQ_PARAM        = "ERR_REQ_PARAM"
	ERR_REQ_PARAM_FORMAT = "ERR_REQ_PARAM_FORMAT"
	ERR_FORMAT           = "ERR_FORMAT"
	ERR_ORDER_REPEAT     = "ERR_ORDER_REPEAT"
	ERR_ORDER            = "ERR_ORDER"
	ERR_ORDER_SEND       = "ERR_ORDER_SEND"
	ERR_SYSTEM           = "ERR_SYSTEM"
	ERR_IDENTITYID       = "ERR_IDENTITYID"
)

var ErrMap = map[string]Err{
	SUCCESS:              Err{Code: "0", Msg: "成功"},
	ERR_SIGN:             Err{Code: "1001", Msg: "验签失败"},
	ERR_REQ_PARAM:        Err{Code: "1002", Msg: "缺少参数"},
	ERR_REQ_PARAM_FORMAT: Err{Code: "1003", Msg: "参数格式错误或不在正确范围内"},
	ERR_IDENTITYID:       Err{Code: "1004", Msg: "无效的Identity ID"},
	ERR_FORMAT:           Err{Code: "1006", Msg: "错误的编码格式"},
	ERR_ORDER_REPEAT:     Err{Code: "1009", Msg: "重复订单"},
	ERR_ORDER:            Err{Code: "1011", Msg: "订单有误，订单信息错误或者订单不存在"},
	ERR_ORDER_SEND:       Err{Code: "2000", Msg: "订单状态不在已发送虚拟码状态"},
	ERR_SYSTEM:           Err{Code: "9999", Msg: "平台内部错误"},
}
