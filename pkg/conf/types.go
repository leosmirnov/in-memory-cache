package conf

import "time"

type Conf struct {
	App   *App   `mapstructure:"app" yaml:"app"`
	API   *API   `mapstructure:"api" yaml:"api"`
	Cache *Cache `mapstructure:"cache" yaml:"api"`
}

type API struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}

type Cache struct {
	CleanupInterval time.Duration `mapstructure:"cleanupInterval" yaml:"cleanupInterval"`
}

type App struct {
	StopTimeout time.Duration `mapstructure:"stopTimeout" yaml:"stopTimeout"`
}
