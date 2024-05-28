package module

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Module struct {
	//Key      string                                  `json:"key"`
	Autostart bool                         `json:"autostart"`
	Name      string                       `json:"name"`
	Type      string                       `json:"type"`
	Command   string                       `json:"command"`
	Dir       string                       `json:"dir"`
	Params    map[string]map[string]Params `json:"params"`
	Exec      *exec.Cmd                    `json:"-"`
	State     string                       `json:"proc_state"`
	isRunning bool
}

func (m *Module) Load(configPath string) error {
	configBytes, _ := os.ReadFile(configPath + string(os.PathSeparator) + "module.json")
	m.Dir = configPath

	if err := json.Unmarshal(configBytes, &m); err != nil {
		return err
	}

	if len(m.Command) >= 2 && m.Command[0:2] == "./" {
		m.Command = m.Dir + string(os.PathSeparator) + strings.Replace(m.Command, "./", "", 1)
	}

	//if m.Autostart && m.Exec != nil {
	//	// Run module
	//}

	return nil
}

func (m *Module) Start() error {
	if len(m.Command) > 0 && !m.isRunning {
		m.Exec = exec.Command(m.Command)
		m.Exec.Dir = m.Dir
	} else {
		return nil
	}

	m.isRunning = true
	if err := m.Exec.Start(); err != nil {
		m.State = "failed"
		return err
	}

	m.State = "run"

	go func(cmd *exec.Cmd) {
		if err := cmd.Wait(); err == nil {
			m.State = "stopped"
			log.Printf("Module %s stopped", cmd.Path)
		} else {
			m.State = "failed"
			log.Printf("Module %s failed: %s", cmd.Path, err)
		}

		m.isRunning = false
	}(m.Exec)

	return nil
}

func (m *Module) Stop() error {
	if err := m.Exec.Process.Kill(); err != nil {
		return err
	}

	return nil
}

func (m *Module) Json() error {
	var data []byte
	err := json.Unmarshal(data, &m)

	return err
}
