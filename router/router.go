package router

import (
	"github.com/gin-gonic/gin"
	"nvidia-gpu-monitoring/gpu-stats"
)
func RunApi (p string) {

	r := gin.Default()
	r.Use(CORS())

	v1 := r.Group("/v1")
	{
		v1.GET("/gpuinfo",gpu_stats.GetGpuInfo)

	}
	r.Run(":"+p)
}