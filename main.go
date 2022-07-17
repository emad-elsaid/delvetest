package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/go-delve/delve/pkg/gobuild"
	"github.com/go-delve/delve/service/debugger"
)

func main() {
	wd, _ := os.Getwd()
	config := debugger.Config{
		WorkingDir:  path.Join(wd, "server"),
		Backend:     "default",
		Foreground:  true,
		ExecuteKind: debugger.ExecutingExistingFile,
	}

	if err := os.Chdir(config.WorkingDir); err != nil {
		log.Fatal(err)
	}

	binPath := path.Join(config.WorkingDir, "debug")
	if _, err := os.Stat(binPath); err == nil {
		gobuild.Remove(binPath)
	}

	if err := gobuild.GoBuild("debug", config.Packages, config.BuildFlags); err != nil {
		log.Fatal(err)
	}

	d, err := debugger.New(&config, []string{path.Join(config.WorkingDir, "debug")})
	if err != nil {
		log.Fatal(err)
	}

	go d.Target().Continue()
	time.Sleep(time.Second * 5)

	_, err = d.Restart(false, "", false, []string{}, [3]string{}, true)
	if err != nil {
		log.Fatal(err)
	}

	go d.Target().Continue()
	for {
		time.Sleep(time.Second)
	}
}
