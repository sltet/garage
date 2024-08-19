package core

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"regexp"
)

var EnvConfigs *EnvConfig

type EnvConfig struct {
	BaseUrl                 string `mapstructure:"BASE_URL"`
	Environment             string `mapstructure:"ENV"`
	GoogleOauthClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOauthClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOauthRedirectUri  string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URI"`
	SessionKey              string `mapstructure:"SESSION_KEY"`
	MysqlUri                string `mapstructure:"MYSQL_URI"`
}

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *EnvConfig) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("BASE_URL", "http://localhost:8080")
	viper.SetDefault("ENV", "local")
	viper.SetDefault("GOOGLE_OAUTH_CLIENT_ID", "593351455385-53l7a4p8a3sfhjl3gm8mpdtl3tf4a5hp.apps.googleusercontent.com")
	viper.SetDefault("GOOGLE_OAUTH_CLIENT_SECRET", "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")
	viper.SetDefault("GOOGLE_OAUTH_REDIRECT_URI", fmt.Sprintf("%s/auth/google/callback", viper.GetString("BASE_URL")))
	viper.SetDefault("SESSION_KEY", "idxnwsy")
	viper.SetDefault("MYSQL_URI", "avnadmin:AVNS_XXC3taUApAqHGirueob@mysql-garage-14-garage-14.g.aivencloud.com:27893/defaultdb")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Reading environment variables from environment...")
		v := reflect.ValueOf(EnvConfig{})

		for i := 0; i < v.NumField(); i++ {
			fieldTag := v.Type().Field(i).Tag
			re := regexp.MustCompile(`"([^"]+)"`)
			match := re.FindStringSubmatch(string(fieldTag))
			if len(match) > 1 {
				viper.BindEnv(match[1])
			}
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshalling env variables", err)
	}
	return
}
