package handler

const (
	PatternParam = "pattern"
)

const (
	RootRoute                = "/api"
	WordRoute                = "/word"
	AddWordRoute             = WordRoute
	QueryWordRoute           = WordRoute
	QueryWordWithParamsRoute = QueryWordRoute + "/:" + PatternParam
)

type Status struct {
	Succeed bool   `json:"succeed"`
	Error   string `json:"error"`
}
