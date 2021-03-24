package util

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func init(){
	dir, _:=os.Getwd()
	dotEnvPath := path.Join(dir, "../../.env")
	fmt.Print(dotEnvPath)
	err := godotenv.Load(dotEnvPath)
	if err != nil {
		log.Fatalf("Error loading dotenv file")
	}
}

