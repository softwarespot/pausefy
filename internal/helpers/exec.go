package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(args []string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing %q: %w", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(out)), nil
}
