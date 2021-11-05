package cons

const (
	FORMAT_TIME = "2006-01-02 15:04:05"
	FORMAT_DATE = "2006-01-02"

	FLAG_QRCODE          = "QRCode:Login:%s"
	FLAG_BANTOKEN_DEVV2  = "BanTokenDev:%s:%s"
	FLAG_BANTOKEN_DEVV2S = "BanTokenDev:%s:%s:%s"

	FLAG_DEVSSOV2 = "DevSSOV2:%s:%s" //牛方案2.0缓存flag
	FLAG_DEVSSO   = "DevSSO:%d:%s"   //牛方案1.0缓存flag

	FLAG_BANNER = "global:banner"

	FLAG_SINGERSONG = "global:singer:%s:song"

	FLAG_KG_SCAN_CODE_LOGIN = "ScanCodeLogin:%s" //酷狗扫码登录的标志

	FREEDEVUSE = "FREEDEVUSE:%s" //车机使用缓存标志
)
