package module

type Params struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Placeholder string `json:"placeholder"`
	Validation  string `json:"validation"`
	Value       any    `json:"value"`
}
