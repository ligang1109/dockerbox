package main

import (
	"dbox/cmd"
	"dbox/dconf"
	"dbox/errno"

	"github.com/goinbox/golog"
	"github.com/goinbox/gomisc"

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
	if len(fargs) == 0 {
		logger.Error([]byte("do not has cmd arg"))
		os.Exit(errno.E_DBOX_INVALID_CMD)
	}

	name := strings.TrimSpace(fargs[0])
	command := cmd.NewCommandByName(name)
	if command == nil {
		logger.Error([]byte("unknown cmd: " + name))
		os.Exit(errno.E_DBOX_INVALID_CMD)
	}

	command.Run(fargs[1:], logger)
}
