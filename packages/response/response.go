package response

import "YudoleChatServer/packages/module"

type GetModulesIdResponse struct {
	Payload *module.Module `json:"payload"`
}

type GetModulesResponse struct {
	Payload map[string]*module.Module `json:"payload"`
}
