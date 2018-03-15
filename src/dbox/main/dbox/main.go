package main

import (
	"dbox/errno"
	"dbox/dconf"

	"github.com/goinbox/golog"
	"github.com/goinbox/gomisc"
	"github.com/goinbox/shell"

	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var logLevel int
	var dconfDir string

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.IntVar(&logLevel, "logLevel", golog.LEVEL_INFO, "global log level")
	fs.StringVar(&dconfDir, "dconfDir", "", "dbox conf dir")
	fs.Parse(os.Args[1:])

	logger, err := golog.NewSimpleLogger(golog.NewStdoutWriter(), logLevel, golog.NewConsoleFormater())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(errno.E_SYS_INIT_LOG_FAIL)
	}

	dconfDir = strings.TrimRight(dconfDir, "/")
	if dconfDir == "" {
		logger.Error([]byte("missing flag dconfDir"))
		flag.PrintDefaults()
		os.Exit(errno.E_SYS_INVALID_RCONF)
	}
	if !gomisc.DirExist(dconfDir) {
		logger.Error([]byte("dconfDir not exist: " + dconfDir))
		os.Exit(errno.E_SYS_INVALID_RCONF)
	}
	logger.Debug([]byte("dconfDir: " + dconfDir))

	dboxConf, err := initRiggerConf(fs, dconfDir, logger)
	if err != nil {
		logger.Error([]byte("initRiggerConf error: " + err.Error()))
		os.Exit(errno.E_SYS_INVALID_RCONF)
	}

	err = genConfByTpl(dboxConf, logger)
	if err != nil {
		logger.Error([]byte("genConfByTpl error: " + err.Error()))
		os.Exit(errno.E_SYS_INVALID_RCONF)
	}

	runAction(dboxConf, logger)
}

func initRiggerConf(fs *flag.FlagSet, dconfDir string, logger golog.ILogger) (*dconf.RiggerConf, error) {
	extArgs := parseExtArgs(fs.Args())
	dboxConf, err := dconf.NewRiggerConf(dconfDir, extArgs, logger)
	if err != nil {
		return nil, err
	}

	err = dboxConf.Parse()
	if err != nil {
		return nil, err
	}

	return dboxConf, nil
}

func parseExtArgs(args []string) map[string]string {
	result := make(map[string]string)

	for _, str := range args {
		item := strings.Split(str, "=")
		if len(item) == 2 {
			result[item[0]] = item[1]
		}
	}

	return result
}

func genConfByTpl(dboxConf *dconf.RiggerConf, logger golog.ILogger) error {
	for key, item := range dboxConf.TplConfItem.Tpls {
		if !gomisc.FileExist(item.Tpl) {
			return errors.New("Gen conf " + key + " tpl " + item.Tpl + " not exists")
		}

		logger.Debug([]byte("gen tpl: " + key))
		logger.Debug([]byte("read file " + item.Tpl))
		tplBytes, err := ioutil.ReadFile(item.Tpl)
		if err != nil {
			return err
		}

		dstString, delay, err := dboxConf.VarConfItem.ParseValueByDefined(string(tplBytes))
		if delay {
			err = errors.New("must not delay")
		}
		if err != nil {
			return err
		}

		logger.Debug([]byte("write dst file " + item.Dst))
		err = ioutil.WriteFile(item.Dst, []byte(dstString), 0644)
		if err != nil {
			return err
		}

		if item.Ln != "" {
			cmd := ""
			cmdPrefix := ""
			if item.Sudo {
				cmdPrefix += "sudo "
			}
			cmd += cmdPrefix + "rm -f " + item.Ln + "; "
			cmd += cmdPrefix + "ln -s " + item.Dst + " " + item.Ln

			shell.RunCmdBindTerminal(cmd)
		}
	}

	return nil
}

func runAction(dboxConf *dconf.RiggerConf, logger golog.ILogger) {
	for _, item := range dboxConf.ActionConfItem.Mkdir {
		cmd := ""
		cmdPrefix := ""
		if item.Sudo {
			cmdPrefix += "sudo "
		}
		if !gomisc.DirExist(item.Dir) {
			cmd += cmdPrefix + "mkdir -p " + item.Dir + "; "
		}
		cmd += cmdPrefix + "chmod " + item.Mask + " " + item.Dir

		logger.Debug([]byte("mkdir run cmd: " + cmd))
		shell.RunCmdBindTerminal(cmd)
	}

	for _, cmd := range dboxConf.ActionConfItem.Exec {
		logger.Debug([]byte("exec run cmd: " + cmd))
		shell.RunCmdBindTerminal(cmd)
	}
}
