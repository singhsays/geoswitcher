package main

import (
	"fmt"
	"os/exec"
)

func switchConfig(config, host, path string) error {
	var (
		cmd  *exec.Cmd
		dest = fmt.Sprintf("%s:%s", host, path)
		err  error
	)
	cmd = exec.Command("scp", config, dest)
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("ssh", host, "service restart_dnsmasq")
	return cmd.Run()
}
