package models

type Config struct {
	PrivateKeyPath string `json:"rsa_private"`
	PublicKeyPath  string `json:"rsa_public"`
}
