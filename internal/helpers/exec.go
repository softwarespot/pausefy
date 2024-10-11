package helpers

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func ExecCmd(args []string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return "", errors.Wrapf(err, "error executing %q", strings.Join(args, " "))
	}
	return string(out), nil
}
