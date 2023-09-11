package utils

import "github.com/spf13/viper"

type Envs struct{
	MongoURI string `mapstructure:"MONGODB_URL"`
}

func LoadEnv(path string)(env Envs, err error ){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	
	viper.AutomaticEnv()
	
	err = viper.ReadInConfig(); 
	if err !=nil {
		return
	}
	
	err = viper.Unmarshal(&env)
	if err !=nil {
		return
	}
	
	return
}