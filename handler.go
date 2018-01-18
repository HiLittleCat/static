package static

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/HiLittleCat/core"
)

const (
	// DefaultMaxAge provides a default caching value of 1 hour.
	DefaultMaxAge = 1 * time.Hour

	assetsDir = "static"
)

var (
	fs = http.FileServer(http.Dir(assetsDir))
)

// Use adds the handler to the default handlers stack.
// Argument maxAge is expressed in seconds and applies to all content (in a production environment only).
func Use(maxAge time.Duration) {
	maxAgeString := fmt.Sprintf("%.f", maxAge.Seconds())
	core.Use(func(c *core.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/"+assetsDir) {
			if core.Production {
				c.ResponseWriter.Header().Set("Cache-Control", "public, max-age="+maxAgeString)
			}
			http.StripPrefix("/"+assetsDir, fs).ServeHTTP(c.ResponseWriter, c.Request)
		} else {
			c.Next()
		}
	})
}
