package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/port-agent/g"
	"github.com/port-agent/funcs"
	"github.com/port-agent/cron"
	"github.com/port-agent/http"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()
	funcs.BuildMappers()
	cron.Collect()
	go http.Start()

	select {}

}



