package test

type Node struct {
	Value   interface{}            `json:"value,omitempty"`
	Type    string                 `json:"type"`
	Details map[string]interface{} `json:"details,omitempty"`
	Nodes   []*Node                `json:"nodes,omitempty"`
}
