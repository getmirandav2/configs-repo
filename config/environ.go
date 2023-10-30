package config

import (
	"strings"
)

type Environment string

const (
	Dev  = Environment("dev")
	Beta = Environment("beta")
	Prod = Environment("prod")
)

func toEnvironment(scope string) Environment {
	normalizedScope := strings.ToLower(scope)
	parts := strings.Split(normalizedScope, "-")

	if strings.EqualFold(normalizedScope, "local") {
		return Dev
	}

	// TODO: Remove this when we have a better way to identify the environment
	if strings.Contains(normalizedScope, "beta") {
		return Beta
	}

	// TODO: Remove this when we have a better way to identify the environment
	if strings.Contains(normalizedScope, "prod") {
		return Prod
	}

	// TODO: Remove this when we have a better way to identify the environment
	if strings.Contains(normalizedScope, "master") {
		return Prod
	}

	return Environment(parts[0])
}
