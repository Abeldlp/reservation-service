package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var Paths = map[string]string{
	"reservations": os.Getenv("RESERVATION_SERVICE"),
	"auth":         os.Getenv("AUTH_SERVICE"),
}

var UnauthorizedPaths = []string{
	"auth",
}

func ProxyRequestToServer(c *gin.Context) {
	requestPath := c.Request.URL.Path
	var service []string = strings.Split(requestPath, "/")

	path := Paths[service[1]] + "/" + strings.Join(service[2:], "/")

	remote, err := url.Parse(path)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = path
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
