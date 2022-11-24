package controller

type Context interface {
	BindJSON(interface{}) error
	JSON(int, interface{})
	GetQuery(key string) (string, bool)
}
