package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type hashicorpTerraformCloud struct {
	Token      string   `yaml:"token"`
	OrgName    string   `yaml:"orgName"`
	Workspaces []string `yaml:"workspaces"`
}

type config struct {
	HashicorpTerraformCloud hashicorpTerraformCloud `yaml:"hashicorpTerraformCloud"`
}

func (c *config) read() error {
	userHome, err := os.UserHomeDir()

	configHierarchy := []string{
		path.Join(userHome, ".config", "inbox-zero", "config.yaml"),
		path.Join(".", "config.yaml"),
	}

	for _, path := range configHierarchy {
		log.Debugf("Looking for config file '%s'", path)

		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			d := yaml.NewDecoder(file)

			if err := d.Decode(&c); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("Could not find config in any of: %s", strings.Join(configHierarchy, ", "))
}
