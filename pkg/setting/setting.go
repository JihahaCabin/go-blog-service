package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	err := vp.ReadInConfig()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
