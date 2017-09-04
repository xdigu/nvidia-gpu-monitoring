package gpu_stats

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"log"
	"encoding/xml"
)

type Query struct {
	DriveVersion string `xml:"driver_version"`
	Gpu  Gpu `xml:"gpu"`
}

type Gpu struct{
	GpuName  string `xml:"product_name"`
	BiosVersion string `xml:"vbios_version"`
	FanSpeed string `xml:"fan_speed"`
	Utilization GpuUtilization `xml:"utilization"`
	Temperature GpuTemperature `xml:"temperature"`
	Clock GpuClock `xml:"clocks"`
}

type GpuUtilization struct {
	GpuUsage string `xml:"gpu_util"`
	MemoryUsage string `xml:"memory_util"`
}

type GpuTemperature struct {
	GpuTemp string `xml:"gpu_temp"`
	}

type GpuClock struct {
	GpuClock string `xml:"graphics_clock"`
	GpuMemClock string `xml:"mem_clock"`
}

func GetGpuInfo(c *gin.Context){
	out, err := exec.Command("nvidia-smi", "-q", "-x").Output()

	

	if err != nil {
		log.Println(err)
	}else {
		var gpu Query
		err = xml.Unmarshal(out, &gpu)
		format := c.Query("format")

		if format == "xml"{
			c.XML(200,gpu)
		} else {
			c.JSON(200,gpu)
		}
	}

}