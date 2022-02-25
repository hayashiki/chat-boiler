package config

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2/google"
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

func GetProject() (string, error) {
	ctx := context.Background()
	cred, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to find default credentials: %w", err)
	}

	if cred.ProjectID == "" {
		return "", fmt.Errorf("project ID not found")
	}

	return cred.ProjectID, nil
}


// GetProject on Google Cloud
func GetProject2() string {
	var (
		project string
		err     error
	)

	log.Println("sokutei")
	if os.Getenv("DATASTORE_EMULATOR_HOST") != "" {
	 	return os.Getenv("DATASTORE_PROJECT_ID")
	}

	project, err = metadata.ProjectID()
	log.Println("sokutei end", project)
	if err != nil {
		if project = os.Getenv("GCP_PROJECT"); project == "" {
			// TODO: if emulator
			project = "local"
		}
	}

	log.Println(project)
	return project
}
