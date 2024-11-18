package types

type PostmanInfo struct {
	PostmanId  string `json:"_postman_id,omitempty"`
	Name       string `json:"name"`
	Schema     string `json:"schema,omitempty"`
	ExporterId string `json:"_exporter_id,omitempty"`
}

type PostmanCollection struct {
	Info PostmanInfo `json:"info"`
	Item []Item      `json:"item"`
}

type Item struct {
	Name     string   `json:"name,omitempty"`
	Request  *Request `json:"request,omitempty"`
	Response []string `json:"response,omitempty"`
	Item     []Item   `json:"item,omitempty"` // For folders containing more items
}

type Request struct {
	Method string   `json:"method"`
	Header []Header `json:"header"`
	Body   Body     `json:"body,omitempty"`
	URL    URL      `json:"url"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Body struct {
	Mode    string      `json:"mode,omitempty"`
	Raw     string      `json:"raw,omitempty"`
	Options BodyOptions `json:"options,omitempty"`
}

type BodyOptions struct {
	Raw RawOptions `json:"raw,omitempty"`
}

type RawOptions struct {
	Language string `json:"language,omitempty"`
}

type URL struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Port     string   `json:"port"`
	Path     []string `json:"path"`
	Query    []Query  `json:"query"`
}

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
