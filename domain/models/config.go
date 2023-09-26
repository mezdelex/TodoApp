package models

type Config struct {
	PrivateKey string `json:"ed25519:private"`
	PublicKey  string `json:"ed25519:public"`
}
