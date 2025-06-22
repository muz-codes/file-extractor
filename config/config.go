package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	//os.Setenv("APY_ENV", "local")
	if os.Getenv("APY_ENV") == "local" {
		err1 := godotenv.Load("/app/.env")
		err2 := godotenv.Load("./../.env")
		if err1 != nil && err2 != nil {
			log.Fatalln(fmt.Printf("error while loading env. env not found at %s and %s", err1.Error(), err2.Error()))
		}
	}
	viper.Set("azure.ai_service_key", os.Getenv("AZURE_AI_SERVICE_KEY"))
	viper.Set("azure.ai_service_endpoint", os.Getenv("AZURE_AI_SERVICE_ENDPOINT"))
	viper.Set("azure.ai_service_region", os.Getenv("AZURE_AI_SERVICE_REGION"))
}

func GetConfig(key string) string {
	return viper.GetString(key)
}
