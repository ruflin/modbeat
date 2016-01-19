// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type ModbeatConfig struct {
	Modbeat struct {
		Paths []PathConfig
	}
}

type PathConfig struct {
	Glob   string `yaml:"glob"`
	Period string `yaml:"period"`
}
