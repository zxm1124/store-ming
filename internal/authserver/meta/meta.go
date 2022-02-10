package meta

type SignInfo struct {
	// Secret 签发密钥
	Secret string
	// Timeout 有效期 hour
	Timeout int
	// Issuer 签发者 "iamctl"
	Issuer string
	// Audience 接收者 "iam.authz.marmotedu.com",
	Audience string
}

type Auth struct {
	// Http对外端口
	ParseHttpPort int
	// 解析token路由
	ParsePath string

	// rpc对外端口
	SignRpcPort int
	// 签发配置
	SignInfo SignSetting
}
type SignSetting struct {
	// Secret 签发密钥
	Secret string
	// Timeout 有效期 hour
	Timeout int
	// Issuer 签发者 "iamctl"
	Issuer string
	// Audience 接收者 "iam.authz.marmotedu.com",
	Audience string
}
