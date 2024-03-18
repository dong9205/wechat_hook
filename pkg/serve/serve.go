package serve

import (
	"net/http"

	"github.com/dong9205/wechat_hook/pkg/logger"
	"github.com/dong9205/wechat_hook/pkg/models"
	"github.com/dong9205/wechat_hook/pkg/wx"
	ginlog "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func Serve() error {
	r := gin.New()

	// Global middleware
	gin.SetMode(viper.GetString("api.mode"))
	r.Use(ginlog.SetLogger(
		ginlog.WithSkipPath([]string{
			viper.GetString("api.health_uri"),
			viper.GetString("api.metric_uri"),
		}),
	))
	r.Use(gin.Recovery())
	r.POST(viper.GetString("api.push_msg"), pushHandler)
	r.GET(viper.GetString("api.health_uri"), metricsHandler)
	r.GET(viper.GetString("api.metric_uri"), heartbeatHandler)
	r.HEAD(viper.GetString("api.metric_uri"), heartbeatHandler)
	return r.Run(viper.GetString("api.listen"))
}

func pushHandler(c *gin.Context) {
	logger := logger.GetLogger()

	pushMsgReq := models.PussMsgRequest{}
	if err := c.ShouldBindJSON(&pushMsgReq); err != nil {
		logger.Error().Err(err).Msg("invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info().Str("dest", pushMsgReq.Dest).Str("msg", pushMsgReq.Msg).Str("dest_type", pushMsgReq.DestType).Msg("receive msg")
	err := wx.GetBotMgr().Send(pushMsgReq.Msg, pushMsgReq.Dest, pushMsgReq.DestType)
	if err != nil {
		logger.Error().Err(err).Msg("send msg failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func heartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
