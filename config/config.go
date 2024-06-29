package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type Apollo struct {
	MetaAddr        string   `yaml:"meta_addr" env:"apollo_addr"`
	AppId           string   `yaml:"app_id" env:"apollo_app_id"`
	Cluster         string   `yaml:"cluster" env:"apollo_cluster"`
	NamespaceNames  []string `yaml:"namespace_names" env:"apollo_namespace"`
	AccesskeySecret string   `yaml:"accesskey_secret" env:"apollo_secret"`
}

func (a *Apollo) IsEmpty() bool {
	return a.MetaAddr == "" ||
		a.AppId == "" ||
		a.Cluster == "" ||
		a.NamespaceNames == nil ||
		a.AccesskeySecret == "" ||
		len(a.NamespaceNames) == 0
}

type ViperConfig struct {
	ConfigPath string `yaml:"path" env:"config_path"`
}

func (v *ViperConfig) IsEmpty() bool {
	return v.ConfigPath == ""
}

func ReadEnvConfig[T any]() *T {
	var a = new(T)
	err := cleanenv.ReadEnv(a)
	if err != nil {
		panic(err)
	}
	return a
}

func Get[T any](namespace string) (*T, error) {
	t, err0 := ApolloGet[T](namespace)
	if err0 == nil {
		return t, nil
	}

	t, err1 := ViperGetAll[T]()
	if err1 == nil {
		return t, nil
	}

	return nil, errors.Join(err0, err1)
}

func MustGet[T any](namespace string) *T {
	if t, err := Get[T](namespace); err != nil {
		panic(err)
	} else {
		return t
	}
}
