package csrf_cookie

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// The test code was created with reference to https://github.com/utrack/gin-csrf/blob/master/csrf_test.go.

func init() {
	gin.SetMode(gin.TestMode)
}

func newServer(options Options) *gin.Engine {
	g := gin.New()

	store := cookie.NewStore([]byte("secret123"))

	g.Use(sessions.Sessions("my_session", store))
	g.Use(Middleware(options))

	return g
}

type requestOptions struct {
	Method  string
	URL     string
	Headers map[string]string
}

func request(server *gin.Engine, options requestOptions) *httptest.ResponseRecorder {
	if options.Method == "" {
		options.Method = "GET"
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(options.Method, options.URL, nil)

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	server.ServeHTTP(w, req)

	if err != nil {
		panic(err)
	}

	return w
}

func TestCookie(t *testing.T) {
	g := newServer(Options{
		Secret: "secret123",
	})

	g.GET("/login", func(c *gin.Context) {
		LoadToken(c, "/", "", false)
	})

	g.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r1 := request(g, requestOptions{URL: "/login"})
	nextCookie := ""
	for _, cookie := range r1.Result().Cookies() {
		nextCookie += cookie.String() + "; "
	}
	r2 := request(g, requestOptions{
		Method: "POST",
		URL:    "/login",
		Headers: map[string]string{
			"Cookie": nextCookie,
		},
	})

	if body := r2.Body.String(); body != "OK" {
		t.Error("Response is not OK: ", body)
	}
}
