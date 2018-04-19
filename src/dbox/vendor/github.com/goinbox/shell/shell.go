/**
* @file shell.go
* @brief tool for exec shell cmd
* @author ligang
* @date 2016-01-28
 */

package shell

import (
	"github.com/goinbox/gomisc"

	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

type ShellResult struct {
	Ok     bool
	Output []byte
}

func NewCmd(cmdStr string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", cmdStr)
}

func RunCmd(cmdStr string) *ShellResult {
	result := &ShellResult{
		Ok: true,
	}

	var err error

	cmd := NewCmd(cmdStr)
	result.Output, err = cmd.CombinedOutput()

	if err != nil {
		result.Ok = false
		result.Output = gomisc.AppendBytes(result.Output, []byte(err.Error()))
	}
	return result
}

func RunCmdBindTerminal(cmdStr string) {
	cmd := NewCmd(cmdStr)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RunAsUser(cmdStr string, username string) *ShellResult {
	curUser, _ := user.Current()
	if username != "" && username != curUser.Username {
		if curUser.Username == "root" {
			cmdStr = fmt.Sprintf(
				"/sbin/runuser %s -c \"%s\"",
				username,
				strings.Replace(cmdStr, "\"", "\\\"", -1),
			)
		} else {
			cmdStr = fmt.Sprintf(
				"sudo -u %s %s",
				username,
				cmdStr,
			)
		}
	}

	return RunCmd(cmdStr)
}

func Rsync(sou string, dst string, excludeFrom string, sshUser string, timeout int) *ShellResult {
	rsyncCmd := MakeRsyncCmd(sou, dst, excludeFrom, timeout)

	return RunAsUser(rsyncCmd, sshUser)
}

func MakeRsyncCmd(sou string, dst string, excludeFrom string, timeout int) string {
	to := strconv.Itoa(timeout)
	rsyncCmd := "/usr/bin/rsync -av -e 'ssh -o StrictHostKeyChecking=no -o ConnectTimeout=" + to + "' --timeout=" + to + " --update "
	_, err := os.Stat(excludeFrom)
	if nil == err {
		rsyncCmd += "--exclude-from='" + excludeFrom + "' "
	}
	rsyncCmd += sou + " " + dst + " 2>&1"

	return rsyncCmd
}

//keyMap := map[string]string{
//	 "param_key" : "shell_key"
//}
func GetParamsFromShell(shell string, keyMap map[string]string) map[string]string {
	type psk struct {
		paramKey string
		shellKey string
	}

	pskl := make([]*psk, len(keyMap))
	i := 0
	for pk, sk := range keyMap {
		pskl[i] = &psk{pk, sk}
		i++
	}

	cmd := "source " + shell + "; "
	for _, v := range pskl {
		cmd += "echo $" + v.shellKey + "; "
	}

	params := map[string]string{}
	result := RunCmd(cmd)
	if !result.Ok {
		return params
	}

	output := strings.Split(string(result.Output), "\n")
	j := 0
	for _, v := range pskl {
		params[v.paramKey] = output[j]
		j++
	}

	return params
}
