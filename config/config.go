package config

import (
	"os"
	"sync"
)

const (
	_defaultScopeEnvironment = "local"
	_defaultConfigPath       = "internal/config/latest"
	_baseDirEnvVar           = "CONFIG_DIR"
)

type Config interface {
	Refresh() error
	Get() *ConfigValues
}

type configImpl struct {
	cv    ConfigValues
	mutex sync.RWMutex
}

type ConfigValues struct {
	Scope       string
	Environment Environment
	Profiles    Profiles
}

var (
	onceConfig   sync.Once
	configValues *ConfigValues
)

func New() (Config, error) {
	var err error
	onceConfig.Do(func() {
		configValues, err = newConfigValues()
	})

	return &configImpl{
		cv: *configValues,
	}, err
}

func (c *configImpl) Refresh() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	configValues, err := newConfigValues()
	if err != nil {
		return err
	}

	c.cv = *configValues

	return nil
}

func (c *configImpl) Get() *ConfigValues {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return &c.cv
}

func Environ() Environment {
	return configValues.Environment
}

func Scope() string {
	return configValues.Scope
}

func getScopeFromEnv() string {
	scope := os.Getenv("SCOPE")
	if scope == "" {
		scope = _defaultScopeEnvironment
	}

	return scope
}

func newConfigValues() (*ConfigValues, error) {
	scope := getScopeFromEnv()
	environment := toEnvironment(scope)

	profiles, err := readProfiles(environment)
	if err != nil {
		return nil, err
	}

	configValues := ConfigValues{
		Scope:       scope,
		Environment: environment,
		Profiles:    *profiles,
	}

	return &configValues, nil
}
