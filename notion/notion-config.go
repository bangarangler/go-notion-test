package notion

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

var (
	Notion_Key  = goDotEnvVar("NOTION_KEY")
	Notion_DB   = goDotEnvVar("NOTION_DB")
	Notion_Page = goDotEnvVar("NOTION_PAGE")
)
