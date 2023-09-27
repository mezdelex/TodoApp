package models

type Config struct {
	PrivateKeyPath string `json:"ed25519:private"`
	PublicKeyPath  string `json:"ed25519:public"`
}
