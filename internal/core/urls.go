package core

import (
	"fmt"
	"net/http"
)

type Resource struct {
	Method           string
	Resource         string
	ShortDescription string
}

func (e *Resource) Url() string {
	return fmt.Sprintf(
		"%s %s",
		e.Method,
		e.Resource,
	)
}

const V202501 = "/v2025.01"
const UserUrl = "/core/users"

var UserPost = Resource{
	Method:   http.MethodPost,
	Resource: UserUrl,
}

var UserPut = Resource{
	Method:   http.MethodPut,
	Resource: UserUrl,
}

var UserGet = Resource{
	Method:   http.MethodGet,
	Resource: UserUrl,
}

var UserDelete = Resource{
	Method:   http.MethodDelete,
	Resource: UserUrl,
}
