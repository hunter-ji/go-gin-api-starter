// @Title ip_limit_rate.go
// @Description
// @Author Hunter 2024/9/4 10:36

package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"go-gin-api-starter/config"
	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/pkg/util/response"
)

// limitIP
// @Description: limit the access frequency of each ip per second, intercept it if the access frequency exceeds the set threshold
// @param perSecondCount every ip is allowed to access the number of times per second
// @return gin.HandlerFunc
func limitIP(perSecondCount int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := database.RDB
		limiter := redis_rate.NewLimiter(rdb)

		res, err := limiter.Allow(
			c,
			fmt.Sprintf("%s-IPLimitRate:%s", config.CommonSplicePrefix, c.ClientIP()),
			redis_rate.PerSecond(perSecondCount),
		)
		if err != nil {
			fmt.Println("rate limit error : ", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    response.StatusErr,
				"message": "failed to limit ip rate",
			})
			return
		}

		if res.Allowed == 0 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    response.StatusErr,
				"message": "too frequent operation",
			})
			return
		}

		c.Next()
	}
}
