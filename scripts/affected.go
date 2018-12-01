package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	filesValue := flag.String("files", "", "files")
	flag.Parse()

	files := strings.Split(*filesValue, ",")
	goModules := make(map[string]struct{}, len(files))

	for _, file := range files {
		if file == "" {
			continue
		}
		env, err := GetEnv(filepath.Dir(file))
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := goModules[env.GoMod]; ok {
			continue
		}
		goModules[env.GoMod] = struct{}{}
		log.Print(filepath.Dir(env.GoMod))
	}
}

type Env struct {
	GoMod string `json:"GOMOD"`
}

func GetEnv(dir string) (*Env, error) {
	buf := &bytes.Buffer{}
	cmd := exec.Command("go", "env", "-json")
	cmd.Dir = dir
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error while running go env command: %v", err)
	}

	var env Env
	if err := json.NewDecoder(buf).Decode(&env); err != nil {
		return nil, fmt.Errorf("error while parsing json: %v", err)
	}

	if env.GoMod == "" {
		return nil, errors.New("no go module")
	}

	return &env, nil
}
