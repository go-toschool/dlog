package dlog

// Message ...
type Message struct {
	Level   string `json:"level,omitempty"`
	Service string `json:"service,omitempty"`
	Error   string `json:"error,omitempty"`
	Info    string `json:"info,omitempty"`
	Warn    string `json:"warn,omitempty"`
}
