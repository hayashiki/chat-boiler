package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
)

type Config struct {
	//SlackBotToken       string `envconfig:"SLACK_BOT_TOKEN" required:"true"`
	IsDev               bool   `envconfig:"IS_DEV" default:"true"`
}

func NewConfig() (Config, error) {
	env := Config{}
	err := envconfig.Process("", &env)
	return env, err
}

// GetProject on Google Cloud
func GetProject() string {
	var (
		project string
		err     error
	)

	log.Println("sokutei")
	if os.Getenv("DATASTORE_EMULATOR_HOST") != "" {
	 	return os.Getenv("DATASTORE_PROJECT_ID")
	}

	project, err = metadata.ProjectID()
	log.Println("sokutei end")
	if err != nil {
		if project = os.Getenv("GCP_PROJECT"); project == "" {
			// TODO: if emulator
			project = "local"
		}
	}

	log.Println(project)
	return project
}
