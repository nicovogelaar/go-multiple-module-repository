package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	goModules := make(map[string]struct{})

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		env, err := GetEnv(filepath.Dir(line))
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := goModules[env.GoMod]; ok {
			continue
		}
		goModules[env.GoMod] = struct{}{}
		fmt.Println(filepath.Dir(env.GoMod))
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
