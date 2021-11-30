package cons

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

var MapMode = map[string]int{
	MODE_DEV:  1,
	MODE_TEST: 1,
	MODE_PROD: 1,
}
