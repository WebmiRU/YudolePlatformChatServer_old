package main

import "os/exec"

type Config struct {
	AutostartModules []string `json:"active_modules"`
}

type Module struct {
	//Key      string                                  `json:"key"`
	Autostart bool                                    `json:"autostart"`
	Name      string                                  `json:"name"`
	Type      string                                  `json:"type"`
	Command   string                                  `json:"command"`
	Path      string                                  `json:"path"`
	Params    map[string]map[string]ModuleConfigParam `json:"params"`
	Exec      *exec.Cmd                               `json:"-"`
	ProcState string                                  `json:"proc_state"`
}

type ModuleConfigParam struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Placeholder string `json:"placeholder"`
	Validation  string `json:"validation"`
}

type GetModulesIdResponse struct {
	Payload *Module `json:"payload"`
}

type GetModulesResponse struct {
	Payload map[string]*Module `json:"payload"`
}
