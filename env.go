package exam

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
)

type EnvLoader struct {
    PatternWindows []string `json:"pattern_windows"`
    PatternDarwin  []string `json:"pattern_darwin"`
    PatternLinux   []string `json:"pattern_linux"`
    PatternDefault []string `json:"pattern_default"`

    Name string `json:"name"`
}

func (el *EnvLoader) Run(source string, args []string) (result string, err error){
    var pattern []string
    fmt.Println("Debug: OS:", runtime.GOOS)
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

    fmt.Println("Debug: Pattern:", pattern)
    srcRuntime := pattern[0]
    comArgs := strings.Join(pattern[1:], " ")
    comArgs = strings.ReplaceAll(comArgs, "{source}", source)
    comArgs = strings.ReplaceAll(comArgs, "{args}", strings.Join(args, " "))

    fmt.Println("Debug: Execute:", srcRuntime, comArgs)

    x := exec.Command(srcRuntime, strings.Split(comArgs, " ")...)
    b, err := x.CombinedOutput()
    fmt.Println("Debug: Output:", string(b))
    if err != nil {
        return
    }

    result = string(b)
    return
}


type EnvManager struct {
    Environments map[string]*EnvLoader
}

func InitializeEnvMan(path string) (*EnvManager, error) {
    fmt.Println("Looking for environments on:", filepath.Join(path, "*.json"))
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

        fmt.Printf("Loaded Enviroment '%v'\n", envl.Name)
        em.Environments[envl.Name] = envl
    }

    return em, nil
}