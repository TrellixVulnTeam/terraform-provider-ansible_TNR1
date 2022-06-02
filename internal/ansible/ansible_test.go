package ansible

import (
	"testing"
)

func TestReadFileFromEmbed(t *testing.T) {
	_, err := ansibleFs.ReadFile("dist/ansiblex")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunAnsibleFromEmbed(t *testing.T) {
	ansible := NewAnsible()
	output, err := ansible.Run("ansiblex", "localhost -m ping")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
}
