package cfg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Cfg - набор переменных окружения необходимых для запуска сервера
type Cfg struct {
	Port   string `mapstructure:"WB_APP_PORT"`
	DbName string `mapstructure:"WB_APP_DBNAME"`
	DbUser string `mapstructure:"WB_APP_DBUSER"`
	DbPass string `mapstructure:"WB_APP_DBPASS"`
	DbHost string `mapstructure:"WB_APP_DBHOST"`
	DbPort string `mapstructure:"WB_APP_DBPORT"`
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("wb_app")
	v.SetConfigFile(".env")

	err := v.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Panic(fmt.Errorf("fatal error config file: %w", err))
	}

	v.SetDefault("WB_APP_PORT", "8080")
	v.SetDefault("WB_APP_DBNAME", "postgres")
	v.SetDefault("WB_APP_DBUSER", "postgres")
	v.SetDefault("WB_APP_DBPASS", "postgres")
	v.SetDefault("WB_APP_DBHOST", "localhost")
	v.SetDefault("WB_APP_DBPORT", "5432")
	v.AutomaticEnv()

	var cfg Cfg

	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
