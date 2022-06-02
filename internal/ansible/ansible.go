package ansible

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:generate rm -rf ./dist
//go:generate cp -r ../../tools/ansible-portable/dist ./dist/
//go:embed dist/*
var ansibleFs embed.FS

type ansible struct {
	tmpDir string
}

func NewAnsible() *ansible {
	return &ansible{
		tmpDir: "/tmp/ansible",
	}
}

func (an *ansible) Run(command string, args ...string) (string, error) {
	if !validAnsibleCommand(command) {
		return "", fmt.Errorf("invalid ansible command: %s", command)
	}

	bin, err := ansibleFs.ReadFile(fmt.Sprintf("dist/%s", command))
	if err != nil {
		return "", err
	}

	pathFile := fmt.Sprintf("%s/%s", an.tmpDir, command)
	err = os.WriteFile(pathFile, bin, 0755)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(pathFile, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func validAnsibleCommand(command string) bool {
	avaliableCommands := map[string]bool{
		"ansiblex":          true,
		"ansible-config":    true,
		"ansible-console":   true,
		"ansible-doc":       true,
		"ansible-galaxy":    true,
		"ansible-inventory": true,
		"ansible-playbook":  true,
		"ansible-pull":      true,
		"ansible-vault":     true,
	}

	_, ok := avaliableCommands[command]
	return ok
}
