package config

// application : struct to hold application level configs
type application struct {
	Name        string `toml:"app_name"`
	BuildMode   string `toml:"build_mode"`
	ListenPort  int    `toml:"listen_port"`
	ListenIP    string `toml:"listen_ip"`
	LogPath     string `toml:"log_path"`
	Environment string
}
