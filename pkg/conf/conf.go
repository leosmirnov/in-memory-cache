package conf

import (
	"os/user"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/leosmirnov/in-memory-cache/pkg/constants"
)

const configFilename = "config"

func Configure(logger logrus.FieldLogger, cfgFilePath string) (*Conf, error) {

	v := viper.New()

	setDefaults(v)
	var homeDir string
	usr, err := user.Current()
	if err == nil {
		homeDir = usr.HomeDir
	}

	v.SetConfigName(configFilename)

	v.AddConfigPath(cfgFilePath)
	v.AddConfigPath(".")
	v.AddConfigPath("./")
	v.AddConfigPath("./contrib/")
	v.AddConfigPath(homeDir)

	if err := v.ReadInConfig(); err != nil {
		if IsConfNotFoundError(err) {
			logger.Warn("no config file found. using default values variables")
			for _, key := range v.AllKeys() {
				val := v.Get(key)
				v.Set(key, val)
			}
		} else {
			return nil, errors.Wrap(err, "unable to read conf file")
		}
	}

	cfg := &Conf{
		API:   &API{},
		Cache: &Cache{},
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config data")
	}

	return cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("api.host", constants.DefaultAPIHost)
	v.SetDefault("api.port", constants.DefaultAPIPort)
	v.SetDefault("cache.cleanupInterval", constants.DefaultCleanupInterval)
	v.SetDefault("app.stopTimeout", constants.DefaultStopTimeout)
}

func IsConfNotFoundError(err error) bool {
	switch err.(type) {
	case viper.ConfigFileNotFoundError:
		return true
	default:
		return false
	}
}
