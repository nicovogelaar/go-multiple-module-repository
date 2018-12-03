package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dirs := make(map[string]struct{})
	goModules := make(map[string]struct{})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dir := filepath.Dir(line)
		if _, ok := dirs[dir]; ok {
			continue
		}
		dirs[dir] = struct{}{}

		goMod, err := getGoMod(dir)
		if err != nil {
			log.Fatal(err)
		}
		if goMod == "" {
			continue
		}
		if _, ok := goModules[goMod]; ok {
			continue
		}
		goModules[goMod] = struct{}{}

		fmt.Println(filepath.Dir(goMod))
	}
}

type env struct {
	GoMod string `json:"GOMOD"`
}

func getGoMod(dir string) (string, error) {
	buf := &bytes.Buffer{}
	cmd := exec.Command("go", "env", "-json")
	cmd.Dir = dir
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error while running go env command: %v", err)
	}

	var env env
	if err := json.NewDecoder(buf).Decode(&env); err != nil {
		return "", fmt.Errorf("error while parsing json: %v", err)
	}

	return env.GoMod, nil
}
