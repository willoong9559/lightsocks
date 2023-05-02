package common

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func NewClientConfig(listenAddr, remoteAddr, password string) *Config {
	return &Config{listenAddr, remoteAddr, password}
}

func NewServerConfig(listenAddr, remoteAddr, password string) *Config {
	return &Config{listenAddr, remoteAddr, NewRandPasswdStr()}
}
