package test_helper

type Req struct {
	Body map[string]interface{}
}
type Expected struct {
	Code int
	Body map[string]interface{}
}
