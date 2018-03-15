package main

import (
	"dbox/cmd"
	"dbox/dconf"
	"dbox/errno"

	"github.com/goinbox/golog"
	"github.com/goinbox/gomisc"

	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var logLevel int
	var dconfPath string

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.IntVar(&logLevel, "logLevel", golog.LEVEL_INFO, "global log level")
	fs.StringVar(&dconfPath, "dconfPath", os.Getenv("HOME")+"/.dconf.json", "dbox conf path")
	fs.Parse(os.Args[1:])

	logger, err := golog.NewSimpleLogger(golog.NewStdoutWriter(), logLevel, golog.NewConsoleFormater())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(errno.E_SYS_INIT_LOG_FAIL)
	}

	dconfPath = strings.TrimRight(dconfPath, "/")
	if dconfPath == "" {
		logger.Error([]byte("missing flag dconfPath"))
		flag.PrintDefaults()
		os.Exit(errno.E_SYS_INVALID_DCONF)
	}
	if !gomisc.FileExist(dconfPath) {
		logger.Error([]byte("dconfPath not exist: " + dconfPath))
		os.Exit(errno.E_SYS_INVALID_DCONF)
	}
	logger.Debug([]byte("dconfPath: " + dconfPath))

	err = dconf.Init(dconfPath)
	if err != nil {
		logger.Error([]byte("dconf init error: " + err.Error()))
		os.Exit(errno.E_SYS_INVALID_DCONF)
	}

	fargs := fs.Args()
	cmd, err := getCmd(fargs)
	if err != nil {
		logger.Error([]byte("get cmd error: " + err.Error()))
		os.Exit(errno.E_DBOX_INVALID_CMD)
	}

	cmd.Run(fargs[1:], logger)
}

func getCmd(fargs []string) (cmd.ICommand, error) {
	if len(fargs) == 0 {
		return nil, errors.New("do not has cmd arg")
	}

	cmdArg := strings.TrimSpace(fargs[0])
	switch cmdArg {
	case "exec":
		return new(cmd.ExecCommand), nil
	case "attach":
		return new(cmd.AttachCommand), nil
	case "start":
		return new(cmd.StartCommand), nil
	case "stop":
		return new(cmd.StopCommand), nil
	case "restart":
		return new(cmd.RestartCommand), nil
	}

	return nil, errors.New("unknown cmd: " + cmdArg)
}
