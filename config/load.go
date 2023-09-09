package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

func Load() Config {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		fmt.Errorf("error while reading from config file: %w", err)
	}

	cfg := Config{}
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error while murshaling confgi: %s", err)
	}

	if cfg.Debug {
		fmt.Printf("%+v\n", cfg)
	}
	return cfg
}
