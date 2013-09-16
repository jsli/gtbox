package sys

import (
	"os/exec"
	"fmt"
)

func ExecCmd(cmd_str string, params []string) (bool, string, error) {
	cmd := exec.Command(cmd_str, params...)
	output, err := cmd.Output()
	if err != nil || !cmd.ProcessState.Success(){
		return false, fmt.Sprintf("%s", output), err
	}
	return true, fmt.Sprintf("%s", output), nil
}