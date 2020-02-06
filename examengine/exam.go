package examengine

import (
	"io/ioutil"
	"path/filepath"
)

type Manager struct {
	env *EnvManager
}

func InitManager(path string) *Manager {
	envman, err := InitializeEnvMan(filepath.Join(path, "envs"))
	if err != nil {
		panic(err)
	}
	return &Manager{env: envman}
}

func (man *Manager) LoadSpec(specPath string) (*Spec, error) {
	sb, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, err
	}

	spec, err := UnmarshalSpec(sb)
	if err != nil {
		return nil, err
	}

	spec.Manager = man
	return &spec, nil
}
