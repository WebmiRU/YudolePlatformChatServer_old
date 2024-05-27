package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"
)

var config Config
var currentDir string
var modules = make(map[string]*Module)

func modulesStateMonitor() {
	fmt.Println("STR")
	for {
		for _, v := range modules {
			if v.Exec == nil {
				continue
			}
		}

		time.Sleep(1 * time.Second)

	}
}

func loadConfig() {
	configBytes, err := os.ReadFile("config.json")

	if err != nil {
		panic("Error while reading 'config.json' file")
	}

	if json.Unmarshal(configBytes, &config) != nil {
		panic("Error while parsing 'config.json' file")
	}

}

func main() {
	loadConfig()

	go modulesStateMonitor()

	currentDir, _ = os.Getwd()
	extList, _ := os.ReadDir("./modules")

	for _, ext := range extList {
		extDir := "./modules/" + ext.Name()
		configBytes, _ := os.ReadFile(extDir + string(os.PathSeparator) + "/config.json")

		var module Module
		json.Unmarshal(configBytes, &module)

		if len(module.Command) >= 2 && module.Command[0:2] == "./" {
			module.Command = currentDir + string(os.PathSeparator) + "modules" + string(os.PathSeparator) + ext.Name() + strings.Replace(module.Command, "./", string(os.PathSeparator), 1)
		}

		if len(module.Command) > 0 {
			module.Exec = exec.Command(module.Command)
		}

		module.Path = currentDir + string(os.PathSeparator) + "modules" + string(os.PathSeparator) + ext.Name()
		module.Autostart = slices.Contains(config.AutostartModules, ext.Name())

		go func(module *Module) {
			if (*module).Exec == nil {
				return
			}

			(*module).ProcState = "run"
			(*module).Exec.Start()
			state, _ := (*module).Exec.Process.Wait()

			if state.ExitCode() == 0 {
				(*module).ProcState = "stopped"
			} else {
				(*module).ProcState = "failed"
			}

		}(&module)

		modules[ext.Name()] = &module
	}

	//for _, v := range modules {
	//	if v.Exec == nil {
	//		continue
	//	}
	//
	//	go func() {
	//		fmt.Println("PROC START")
	//		if err := v.Exec.Start(); err != nil {
	//			log.Fatalln(err)
	//		}
	//
	//		time.Sleep(6 * time.Second)
	//		fmt.Println("PROC STOP")
	//
	//		v.Exec.Process.Kill()
	//	}()
	//}

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/modules", modulesIndexHandler)
	router.HandleFunc("/api/modules/{id}", modulesIdHandler)
	router.HandleFunc("/api/modules/{id}/autostart/{state:[0,1]}", modulesIdSetAutostartHandler)
	http.Handle("/", router)

	http.ListenAndServe(":80", nil)
}
