package dns

import (
	"os/exec"
)

func CmdHost(domain string) (string, error) {
	cmd := exec.Command("host", "-a", domain)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func CmdDig(domain string) (string, error) {
	cmd := exec.Command("dig", domain)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
