package gooselite

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestLiteBinary(t *testing.T) {
	t.Parallel()

	dir, err := ioutil.TempDir("", "tmptest")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)             // clean up
	defer os.Remove("./bin/lite-goose") // clean up

	// this has to be done outside of the loop
	// since go only supports space separated tags list.
	cmd := exec.Command("go", "build", "-tags='no_postgres no_mysql no_sqlite3'", "-o", "./bin/lite-goose", "./cmd/goose")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s:\n%v\n\n%s", err, cmd, out)
	}

	commands := []string{
		fmt.Sprintf("./bin/lite-goose -dir=%s create user_indices sql", dir),
		fmt.Sprintf("./bin/lite-goose -dir=%s fix", dir),
	}

	for _, cmd := range commands {
		args := strings.Split(cmd, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Env = os.Environ()

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s:\n%v\n\n%s", err, cmd, out)
		}
	}
}
