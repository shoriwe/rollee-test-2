package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/rollee-test-2/common/sqlite"
	"github.com/shoriwe/rollee-test-2/controller"
)

type Handler struct {
	c *controller.Controller
}

func NewTest(t *testing.T) (*httpexpect.Expect, func()) {
	engine := New(controller.New(sqlite.NewTest()))
	server := httptest.NewServer(engine)
	return httpexpect.Default(t, server.URL+RootRoute), server.Close
}

func New(c *controller.Controller) *gin.Engine {
	h := &Handler{c: c}
	engine := gin.Default()
	root := engine.Group(RootRoute)
	root.POST(AddWordRoute, h.AddWord)
	root.GET(QueryWordWithParamsRoute, h.QueryWord)
	return engine
}
