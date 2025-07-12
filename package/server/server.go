package server

import (
	"net/http"

	"github.com/blang/semver/v4"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type BumpyServer struct {
	Engine *gin.Engine
}

func New() BumpyServer {
	router := gin.Default()

	return BumpyServer{
		Engine: router,
	}
}

func (s BumpyServer) Run() {
	s.Engine.GET("/major/:version", func(c *gin.Context) {
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

	s.Engine.GET("/minor/:version", func(c *gin.Context) {
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

	s.Engine.GET("/patch/:version", func(c *gin.Context) {
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

	s.Engine.Run(":8080")
}
