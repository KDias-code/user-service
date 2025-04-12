package configs

import "github.com/spf13/viper"

type Configs struct {
	Port       string `mapstructure:"PORT"`
	Db         string `mapstructure:"DB"`
	GmailLogin string `mapstructure:"GMAIL_LOGIN"`
	GmailPass  string `mapstructure:"GMAIL_PASS"`
	GmailHost  string `mapstructure:"GMAIL_HOST"`
	GmailPort  string `mapstructure:"GMAIL_PORT"`
	RedisHost  string `mapstructure:"REDIS_HOST"`
	RedisPort  string `mapstructure:"REDIS_PORT"`
	RedisDB    string `mapstructure:"REDIS_DB"`
	RedisPass  string `mapstructure:"REDIS_PASS"`
}

func LoadConfigs() (*Configs, error) {
	conf := new(Configs)

	v := viper.New()
	v.AutomaticEnv()

	err := v.BindEnv("PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("GMAIL_LOGIN")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("GMAIL_PASS")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("GMAIL_HOST")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("GMAIL_PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_HOST")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_DB")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("REDIS_PASS")
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
