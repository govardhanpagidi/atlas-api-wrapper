package atlasResponse

type AtlasRespone struct {
	Response       interface{} `json:",omitempty"`
	HttpStatusCode int
	HttpError      string `json:",omitempty"`
}
