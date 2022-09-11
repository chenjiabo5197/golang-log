package main

import (
	"flag"
	"fmt"
	logger "github.com/shengkehua/xlog4go"
)

var (
	logFile = flag.String("l", "./conf/log.json", "log config file path")
)

func initLog(path string) (err error) {
	err = logger.SetupLogWithConf(path)
	if err != nil {
		fmt.Printf("log init fail, err=%s\n", err.Error())
		return
	}
	return nil
}

func main() {

	// init log
	if err := initLog(*logFile); err != nil {
		panic(err)
	}
	logger.Info("init log done!")
}
