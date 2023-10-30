package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/melisource/fury_go-toolkit-config/v2/pkg/config"
	"github.com/mercadolibre/fury_rampup-miranda/internal/config/profiles"
)

const (
	logging        = "logging"
	infrastructure = "infrastructure"
	translation    = "translation"
)

type Profiles struct {
	Infrastructure profiles.Infrastructure
	Logging        profiles.Logging
	Translation    profiles.Translation
}

func read[T any](profile string) (*T, error) {
	b, err := config.Read(profile)
	if err != nil {
		return nil, err
	}

	var p T
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func readProfiles(environment Environment) (*Profiles, error) {
	if environment == Dev {
		if err := os.Setenv(_baseDirEnvVar, path.Join(RootDir(), _defaultConfigPath)); err != nil {
			return nil, err
		}
	}

	infrastructure, err := read[profiles.Infrastructure](infrastructure)
	if err != nil {
		return nil, err
	}

	logging, err := read[profiles.Logging](logging)
	if err != nil {
		return nil, err
	}

	translation, err := read[profiles.Translation](translation)
	if err != nil {
		return nil, err
	}

	return &Profiles{
		Infrastructure: *infrastructure,
		Logging:        *logging,
		Translation:    *translation,
	}, nil
}
