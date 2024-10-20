# gin-csrf-cookie
This is an implementation of CSRF protection using cookies with gin-csrf.

This library was created using and with reference to https://github.com/utrack/gin-csrf.

# Installation

```
go get github.com/okoshi-f/gin-csrf-cookie
```

# Usage

```go
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/okoshi-f/gin-csrf-cookie"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{Secret: "secret123"}))

	r.GET("/protected", func(c *gin.Context) {
		csrf.LoadToken(c, "/", "localhost", false)
		c.String(200, "OK")
	})

	r.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	r.Run(":8080")
}
```

# Disclaimer

We make no warranty for any damages caused by the use of this library.
