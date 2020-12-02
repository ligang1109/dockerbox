package main

import (
	"dbox/cmd"
	"dbox/dconf"
	"dbox/errno"

	"github.com/goinbox/golog"
	"github.com/goinbox/gomisc"

	"flag"
	"os"
	"strings"
)

func main() {
	var logLevel int
	var dconfPath string

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.IntVar(&logLevel, "logLevel", golog.LevelInfo, "global log level")
	fs.StringVar(&dconfPath, "dconfPath", os.Getenv("HOME")+"/.dconf.json", "dbox conf path")
	fs.Parse(os.Args[1:])

	logger := golog.NewSimpleLogger(golog.NewConsoleWriter(), golog.NewConsoleFormater(golog.NewSimpleFormater())).
		SetLogLevel(logLevel)

	dconfPath = strings.TrimRight(dconfPath, "/")
	if dconfPath == "" {
		logger.Error([]byte("missing flag dconfPath"))
		flag.PrintDefaults()
		os.Exit(errno.ESysInvalidDconf)
	}
	if !gomisc.FileExist(dconfPath) {
		logger.Error([]byte("dconfPath not exist: " + dconfPath))
		os.Exit(errno.ESysInvalidDconf)
	}
	logger.Debug([]byte("dconfPath: " + dconfPath))

	err := dconf.Init(dconfPath)
	if err != nil {
		logger.Error([]byte("dconf init error: " + err.Error()))
		os.Exit(errno.ESysInvalidDconf)
	}

	fargs := fs.Args()
	if len(fargs) == 0 {
		logger.Error([]byte("do not has cmd arg"))
		os.Exit(errno.EDboxInvalidCmd)
	}

	name := strings.TrimSpace(fargs[0])
	command := cmd.NewCommandByName(name)
	if command == nil {
		logger.Error([]byte("unknown cmd: " + name))
		os.Exit(errno.EDboxInvalidCmd)
	}

	command.Run(fargs[1:], logger)
}
