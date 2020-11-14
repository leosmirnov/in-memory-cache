package conf

import (
	"github.com/spf13/viper"
)

func Configure() {

	v := viper.New()

	setDefaults(v)

}

func setDefaults(v *viper.Viper) {

	v.SetDefault("api.host", constants.DefaultAPIHost)
	v.SetDefault("api.port", constants.DefaultAPIPort)
}
