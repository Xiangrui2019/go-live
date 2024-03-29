package restfulapi

import (
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Serve(l net.Listener) error {
	router := httprouter.New()

	// APP Restful API
	router.POST("/app/:appname/create", CreateAppHandler)
	router.GET("/app", ListAppsHandler)
	router.GET("/app/:appid/get", GetAppByIdHandler)
	router.DELETE("/app/:appid/del", DeleteAppByIdHandler)

	// Live Restful API
	router.POST("/live/:appname/:livename/create", CreateLiveHandler)
	router.GET("/live", ListLivesHandler)
	router.GET("/live/:appname", ListLivesByAppnameHandler)
	router.GET("/live/:appname/:liveid/get", GetLiveByIdHandler)
	router.PUT("/live/:appname/:liveid/refershtoken", RefershLiveTokenByIdHandler)
	router.DELETE("/live/:appname/:liveid/del", DeleteLiveByIdHandler)

	http.Serve(l, router)
	return nil
}
