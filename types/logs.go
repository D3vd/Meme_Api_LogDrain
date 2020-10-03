package types

// RouterLog - Struct for heroku router log line
type RouterLog struct {
	At      string
	Method  string
	Path    string
	Host    string
	Fwd     string
	Service string
	Status  string
}
