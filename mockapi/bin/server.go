package main

import (
	"fmt"
	"net/http"

	"github.com/MonsterYNH/terraform_demo/mockapi"
	"github.com/gin-gonic/gin"
)

func Run(host, port string) error {
	engine := gin.Default()

	engine.GET("/api/v1/vdc/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, mockapi.GetVdcList())
	})

	engine.GET("/api/v1/vdc/instance", func(c *gin.Context) {
		id := c.Query("id")
		vdc, err := mockapi.ReadVdc(id)
		if err != nil {
			c.JSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, vdc)
	})

	engine.GET("/api/v1/vdc/create", func(c *gin.Context) {
		specs := c.Query("specs")
		vdc, err := mockapi.CreateVdc(specs)
		if err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, vdc)
	})

	engine.GET("/api/v1/vdc/update", func(c *gin.Context) {
		id := c.Query("id")
		specs := c.Query("specs")
		vdc, err := mockapi.UpdateVdc(id, specs)
		if err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, vdc)
	})

	engine.GET("/api/v1/vdc/delete", func(c *gin.Context) {
		id := c.Query("id")
		if err := mockapi.DeleteVdc(id); err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, nil)
	})

	engine.GET("/api/v1/vdcspec/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, mockapi.GetVdcSepcList())
	})

	engine.GET("/api/v1/vdcspec/instance", func(c *gin.Context) {
		id := c.Query("id")
		vdcSpec, err := mockapi.ReadVdcSpecs(id)
		if err != nil {
			c.JSON(http.StatusOK, err)
		}
		c.JSON(http.StatusOK, vdcSpec)
	})

	engine.GET("/api/v1/vdcspec/create", func(c *gin.Context) {
		specs := c.Query("specs")
		vdcSpec, err := mockapi.CreateVdcSpecs(specs)
		if err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, vdcSpec)
	})

	engine.GET("/api/v1/vdcspec/update", func(c *gin.Context) {
		id := c.Query("id")
		specs := c.Query("specs")
		vdcSpec, err := mockapi.UpdateVdcSpecs(id, specs)
		if err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, vdcSpec)
	})

	engine.GET("/api/v1/vdcspec/delete", func(c *gin.Context) {
		id := c.Query("id")
		if err := mockapi.DeleteVdcSpecs(id); err != nil {
			c.JSON(http.StatusNotModified, err)
		}
		c.JSON(http.StatusOK, nil)
	})

	return engine.Run(fmt.Sprintf("%s:%s", host, port))
}

func main() {
	if err := Run("0.0.0.0", "8080"); err != nil {
		panic(err)
	}
}
