package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func thenConfigGettersWorkAsExpected(t *testing.T, config *Config) {
	// assert.Equal(t, config.environmentConfig, config.EnvironmentConfig())
	assert.Equal(t, config.scope, config.Scope())
	assert.Equal(t, config.environment, config.Environment())
}

func TestNewConfig_givenLocalScopeOSEnvironmentVariablesWhenCreatingConfigThenReturnsDevEnvironmentConfigs(t *testing.T) {
	oldScope := os.Getenv("SCOPE")
	defer func() { _ = os.Setenv("SCOPE", oldScope) }()

	inputs := []struct {
		message             string
		scopeValue          string
		expectedEnvironment enums.Environment
	}{
		{
			message:             "local environment",
			scopeValue:          "local",
			expectedEnvironment: enums.DevEnvironment,
		},
		{
			message:             "beta environment",
			scopeValue:          "test-scope-beta",
			expectedEnvironment: enums.BetaEnvironment,
		},
		{
			message:             "prod environment",
			scopeValue:          "test-scope",
			expectedEnvironment: enums.ProdEnvironment,
		},
	}
	for _, input := range inputs {
		t.Run(input.message, func(t *testing.T) {
			_ = os.Setenv("SCOPE", input.scopeValue)

			config, _ := NewConfig()

			assert.Equal(t, input.expectedEnvironment, config.environment)
			assert.Equal(t, input.scopeValue, config.scope)
			thenConfigGettersWorkAsExpected(t, config)
		})
	}
}

func TestToEnvironment(t *testing.T) {
	inputs := []struct {
		message             string
		scope               string
		expectedEnvironment enums.Environment
	}{
		{
			message:             "given scope without suffix calling toEnvironment then returns ProdEnvironment",
			scope:               "some-scope",
			expectedEnvironment: enums.ProdEnvironment,
		},
		{
			message:             "given 'prod' scope when calling toEnvironment then returns ProdEnvironment",
			scope:               "prod",
			expectedEnvironment: enums.ProdEnvironment,
		},
		{
			message:             "given 'consumer-prod' scope when calling toEnvironment then returns ProdEnvironment",
			scope:               "consumer-prod",
			expectedEnvironment: enums.ProdEnvironment,
		},
		{
			message:             "given 'beta' scope when calling toEnvironment then returns ProdEnvironment",
			scope:               "beta",
			expectedEnvironment: enums.BetaEnvironment,
		},
		{
			message:             "given 'consumer-beta' scope when calling toEnvironment then returns ProdEnvironment",
			scope:               "consumer-beta",
			expectedEnvironment: enums.BetaEnvironment,
		},
		{
			message:             "given 'local' scope when calling toEnvironment then returns ProdEnvironment",
			scope:               "local",
			expectedEnvironment: enums.DevEnvironment,
		},
	}

	for _, input := range inputs {
		t.Run(input.message, func(t *testing.T) {
			environment := toEnvironment(input.scope)
			assert.Equal(t, input.expectedEnvironment, environment)
		})
	}
}
*/

func TestNew_WhenCalled_ReturnsConfig(t *testing.T) {
	config, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}
