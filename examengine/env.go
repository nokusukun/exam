package examengine

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/shlex"
)

type EnvLoader struct {
	PatternWindows []string `json:"pattern_windows"`
	PatternDarwin  []string `json:"pattern_darwin"`
	PatternLinux   []string `json:"pattern_linux"`
	PatternDefault []string `json:"pattern_default"`

	Name string `json:"name"`
}

func (el *EnvLoader) Run(context context.Context, source string, args []string) (cmd *exec.Cmd, err error) {
	var pattern []string
	Log("Debug: OS:", runtime.GOOS)
	switch runtime.GOOS {
	case "linux":
		pattern = el.PatternLinux
	case "darwin":
		pattern = el.PatternDarwin
	case "windows":
		pattern = el.PatternWindows
	}

	if len(pattern) == 0 {
		pattern = el.PatternDefault
	}

	Log("Debug: Pattern:", pattern)
	srcRuntime := pattern[0]
	comArgs := strings.Join(pattern[1:], " ")
	comArgs = strings.ReplaceAll(comArgs, "{source}", source)
	comArgs = strings.ReplaceAll(comArgs, "{args}", strings.Join(args, " "))

	Log("Debug: Execute:", srcRuntime, comArgs)

	// TODO: Change exec.Command to utilize exec.CommandContext and return the unran *Cmd
	commandArgs, err := shlex.Split(comArgs)
	if err != nil {
		return
	}
	cmd = exec.CommandContext(context, srcRuntime, commandArgs...)
	return
}

type EnvManager struct {
	Environments map[string]*EnvLoader
}

func InitializeEnvMan(path string) (*EnvManager, error) {
	Log("Looking for environments on:", filepath.Join(path, "*.json"))
	envPath, err := filepath.Glob(filepath.Join(path, "*.json"))
	if err != nil {
		return nil, err
	}

	em := &EnvManager{Environments: make(map[string]*EnvLoader)}

	for _, env := range envPath {
		jb, err := ioutil.ReadFile(env)
		if err != nil {
			return nil, err
		}

		envl := &EnvLoader{}
		err = json.Unmarshal(jb, envl)
		if err != nil {
			return nil, err
		}

		Log("Loaded Enviroment '%v'\n", envl.Name)
		em.Environments[envl.Name] = envl
	}

	return em, nil
}
