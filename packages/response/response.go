package response

import "YudoleChatServer/packages/module"

type GetModulesId struct {
	Payload *module.Module `json:"payload"`
}

type GetModules struct {
	Payload map[string]*module.Module `json:"payload"`
}
