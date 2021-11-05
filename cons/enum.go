package cons

const (
	ORDER_SUCC   = "succ"
	ORDER_FAIL   = "failed"
	ORDER_CREATE = "created"
	ORDER_REVERT = "revert"
	ORDER_RETRY  = "retry"
)

//特殊指令
const (
	//查询是否具体该活动参与权限指定单号内容
	ACTIVITY_POWER_SELECT = "SELECT"
)

const (
	ORDER_IOT   = 1 //iot直充
	ORDER_GIVE  = 2 //赠送充值
	ORDER_UNITE = 3 //联合直充
)

const (
	SP_KW = "KW"
	SP_KG = "KG"
	SP_QM = "QM"
)

const (
	MODE_DEV  = "dev"
	MODE_TEST = "test"
	MODE_PROD = "prod"
)

var EnvMap = map[string]string{
	MODE_DEV:  "10.17.4.158",
	MODE_TEST: "commercialjoin.kugou.com",
	MODE_PROD: "commercial.kugou.com",
}

const (
	IDC_GZ = "GZ"
	IDC_BJ = "BJ"
)

var MapMode = map[string]int{
	MODE_DEV:  1,
	MODE_TEST: 1,
	MODE_PROD: 1,
}

const LEN_MD5 = 32

const (
	DTYPE_CLASS_TV         = "tv"
	DTYPE_CLASS_CAR        = "car"
	DTYPE_CLASS_VBOX       = "voicebox"
	DTYPE_CLASS_KSING      = "ksing"
	DTYPE_CLASS_TVOPERATOR = "tvoperator"
	DTYPE_CLASS_APPLE_VBOX = "applebox"

	DTYPE_CLASS_COMMERCIAL  = "commercial"
	DTYPE_CLASS_COMMERCIAL1 = "commercial1"
	DTYPE_CLASS_COPYRIGHT   = "copyright"

	DTYPE_CLASS_CARAPK1 = "carapk1"
	DTYPE_CLASS_CARAPK2 = "carapk2"
	DTYPE_CLASS_CARSDK  = "carsdk"
)

const ROOM_PRI_VIP = "1" // 房间类型1 需要开通vip才可观看
const ROOM_PRI_BUY = "2" // 房间类型2 需要购买才可观看

const ( //支付渠道
	PAYTYPE_ALI    = 1
	PAYTYPE_WECHAT = 2
	PAYTYPE_GIFT   = 3
)

const (
	OSTATUS_PROC   = "process" // 创建订单后的订单初始状态
	OSTATUS_EXPIRE = "expired" // 后台定时任务设置
	OSTATUS_FAIL   = "fail"    // 支付回调告知-支付失败
	OSTATUS_GIVEUP = "giveup"  // 支付回调告知-放弃支付
	OSTATUS_PAY    = "payed"   // 支付回调告知-支付成功
	OSTATUS_SUCC   = "succ"
	OSTATUS_REFUND = "refund"
)

const (
	SECOPEN_MEDIA  = "1" //媒资开放平台接口
	SECOPEN_NEWAPI = "2" //平台新版签名接口
	SECOPEN_OLDAPI = "3" //平台旧版签名接口
	SECOPEN_SING   = "4" //唱唱签名接口
)

// 鉴权过滤方式
const (
	FilterNoCiqu = 1 // 过滤不可唱
	FilterMusic  = 2 // 过滤纯音乐
)
