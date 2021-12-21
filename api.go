package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinFunc struct {
	Path   string
	Method string
	Func   gin.HandlerFunc
}

type HTTPFunc struct {
	Path   string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

// HandleHTTP ...
// @Description:
type HandleHTTP interface {
	ListHTTP() []*HTTPFunc
}

// HandleGin ...
// @Description:
type HandleGin interface {
	ListGin() []*GinFunc
}

type API interface {
	APIValid() bool
	HandleHTTP
	HandleGin
}
