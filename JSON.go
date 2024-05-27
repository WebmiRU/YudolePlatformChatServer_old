package main

type Config struct {
	ActiveModules []string `json:"active_modules"`
}

type Module struct {
	//Key      string                                  `json:"key"`
	IsActive bool                                    `json:"is_active"`
	Name     string                                  `json:"name"`
	Type     string                                  `json:"type"`
	Command  string                                  `json:"command"`
	Path     string                                  `json:"path"`
	Params   map[string]map[string]ModuleConfigParam `json:"params"`
}

type ModuleConfigParam struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	Validation  string `json:"validation"`
}

type GetModulesIdResponse struct {
	Payload Module `json:"payload"`
}

type GetModulesResponse struct {
	Payload map[string]Module `json:"payload"`
}
