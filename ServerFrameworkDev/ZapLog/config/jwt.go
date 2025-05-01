package config

type Jwt struct {
	Issuer     string `json:"issuer"`
	Secret     string `json:"secret"`
	ExpireTime int64  `json:"expireTime"`
}
