package entities

type ResponseFile struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	Path string `json:"path,omitempty"`
}
