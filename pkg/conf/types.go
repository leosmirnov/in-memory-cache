package conf

type Conf struct {
	API *API `mapstructure:"api" yaml:"api"`
}

type API struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}
