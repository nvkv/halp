package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Config struct {
	Datasource DatasourceConfig `hcl:"datasource"`
}

type DatasourceConfig struct {
	GoogleSheets GSheetsConfig `hcl:"google_sheets"`
}

type GSheetsConfig struct {
	SheetID             string `hcl:"sheet_id"`
	CredentialsFilePath string `hcl:"credentials_file"`
	TokenFilePath       string `hcl:"token_file"`
	Range               string `hcl:"range"`
}

func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(&result, hclParseTree); err != nil {
		return nil, err
	}

	return result, nil
}

func LoadDefaultConfig() (Config, error) {
	data, err := ioutil.ReadFile(DefaultConfigLocation())
	if err != nil {
		return Config{}, err
	}

	cfg, err := ParseConfig(string(data))
	return *cfg, err
}

func DefaultConfigLocation() string {
	return fmt.Sprintf("config.hcl")
}
