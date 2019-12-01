package ctrl

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SayHello returns hello
func SayHello(c *gin.Context) {
	c.String(200, fmt.Sprintf("[%s]Hello API\n", time.Now().Format(time.RFC3339)))
}
