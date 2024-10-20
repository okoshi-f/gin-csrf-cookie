package csrf_cookie

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// The option saves the CSRF middleware settings using cookies.
type Options struct {
	Secret        string
	IgnoreMethods []string
	ErrorFunc     gin.HandlerFunc
}

func tokenGetter(c *gin.Context) string {
	token, err := c.Cookie("_csrf")
	if err != nil {
		return ""
	}
	return token
}

func defaultErrorFunc(c *gin.Context) {
	c.String(400, "CSRF token mismatch")
	c.Abort()
}

// Middleware validates CSRF tokens in cookies.
func Middleware(options Options) gin.HandlerFunc {
	if options.ErrorFunc == nil {
		options.ErrorFunc = defaultErrorFunc
	}

	return csrf.Middleware(csrf.Options{
		Secret:        options.Secret,
		IgnoreMethods: options.IgnoreMethods,
		ErrorFunc:     options.ErrorFunc,
		TokenGetter:   tokenGetter,
	})
}

// Set a verification token for CSRF in the cookie.
func LoadToken(c *gin.Context, path string, domain string, secure bool) {
	token := csrf.GetToken(c)

	c.SetCookie("_csrf", token, 0, path, domain, secure, true)
}
