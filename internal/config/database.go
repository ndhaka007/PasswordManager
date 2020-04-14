package config

import (
	"fmt"
)

// Database : struct to hold live / test Database config
type Database struct {
	Dialect            string `toml:"dialect"`
	Protocol           string `toml:"protocol"`
	Host               string `toml:"host"`
	Port               int    `toml:"port"`
	Username           string `env:"username"`
	Password           string `env:"password"`
	Name               string `toml:"name"`
	MaxOpenConnections int    `toml:"max_open_connections"`
	MaxIdleConnections int    `toml:"max_idle_connections"`
}

// SqlConnectionDSNFormat : DSN for connecting mysql
const SqlConnectionDSNFormat = "%s:%s@%s(%s:%d)/%s"

// URL : gives formatted postgresql url.
func (c Database) URL() string {

	// charset=utf8: uses utf8 character set data format
	// parseTime=true: changes the output type of DATE and DATETIME values to time.Time instead of []byte / strings
	// loc=Local: Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location
	return fmt.Sprintf(
		SqlConnectionDSNFormat,
		c.Username,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		c.Name)
}
