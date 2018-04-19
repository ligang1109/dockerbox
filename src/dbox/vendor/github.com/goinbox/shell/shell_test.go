package shell

import (
	"os/user"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	cmd := "ls -l"

	result := RunCmd(cmd)
	if !result.Ok {
		t.Errorf("run command failed, command: %s", cmd)
	}

	if string(result.Output) == "" {
		t.Errorf("run command return empty, command: %s", cmd)
	}
}

func TestRunAsUser(t *testing.T) {
	cmd := "ls -l"
	user := "root"

	result := RunAsUser(cmd, user)
	if !result.Ok {
		t.Errorf("run command as %s failed, command: %s", user, cmd)
	}

	if string(result.Output) == "" {
		t.Errorf("run commandas %s return empty, command: %s", user, cmd)
	}
}

func TestRsync(t *testing.T) {

	sou := "./tmp/rsync/sou/"
	dst := "./tmp/rsync/dst/"
	file := "rsync.txt"

	cmd := "mkdir -p " + sou + "; mkdir -p " + dst + "; /bin/echo 'rsync sou' > " + sou + file
	result := RunCmd(cmd)
	if !result.Ok {
		t.Fatalf("run command failed, command: %s", cmd)
	}

	currentUser, _ := user.Current()
	sshUser := currentUser.Username

	result = Rsync(sou, dst, "", sshUser, 3)
	if !result.Ok {
		t.Errorf("rsync failed")
	}

	if strings.Index(string(result.Output), file) == -1 {
		t.Errorf("rsync file %s failed", file)
	}
}

func TestGetParamsFromShell(t *testing.T) {
	shell := "testdata/params.sh"
	paramMap := map[string]string{
		"user_name": "user_name",
		"nick_name": "nick_name",
		"user_sex":  "sex",
	}
	params := GetParamsFromShell(shell, paramMap)
	if len(params) != len(paramMap) {
		t.Errorf("params count from shell error, expect %d, got %d", len(paramMap), len(params))
	}

	expectUserName := "zhangsan"
	userName := params["user_name"]
	if userName != expectUserName {
		t.Errorf("params userName from shell error, expect %s, got %s", expectUserName, userName)
	}
}
