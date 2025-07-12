package main

import (
	"net/http"

	"github.com/blang/semver/v4"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/major/:version", func(c *gin.Context) {
		inputVersion := c.Param("version")

		v, err := semver.Make(inputVersion)
		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = v.IncrementMajor()

		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"version": v.String(),
		})
	})

	router.GET("/minor/:version", func(c *gin.Context) {
		inputVersion := c.Param("version")

		v, err := semver.Make(inputVersion)
		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = v.IncrementMinor()

		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"version": v.String(),
		})
	})

	router.GET("/patch/:version", func(c *gin.Context) {
		inputVersion := c.Param("version")

		v, err := semver.Make(inputVersion)
		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = v.IncrementPatch()

		if err != nil {
			log.Errorf("%s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"version": v.String(),
		})
	})

	router.Run(":8080")
}
