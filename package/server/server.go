package server

import (
	"fmt"
	"net/http"

	"github.com/blang/semver/v4"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

const (
	v1 = iota + 1
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
	apiV1 := s.Engine.Group(fmt.Sprintf("/api/v%d", v1))
	{
		apiV1.GET("/major/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if len(c.Query("pre-release")) != 0 {
				pVName, err := semver.NewPRVersion(c.Query("pre-release"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Pre = append(v.Pre, pVName)
			}

			if len(c.Query("build")) != 0 {
				bVName, err := semver.NewBuildVersion(c.Query("build"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Build = append(v.Build, bVName)
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

		apiV1.GET("/minor/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if len(c.Query("pre-release")) != 0 {
				pVName, err := semver.NewPRVersion(c.Query("pre-release"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Pre = append(v.Pre, pVName)
			}

			if len(c.Query("build")) != 0 {
				bVName, err := semver.NewBuildVersion(c.Query("build"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Build = append(v.Build, bVName)
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

		apiV1.GET("/patch/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if len(c.Query("pre-release")) != 0 {
				pVName, err := semver.NewPRVersion(c.Query("pre-release"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Pre = append(v.Pre, pVName)
			}

			if len(c.Query("build")) != 0 {
				bVName, err := semver.NewBuildVersion(c.Query("build"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
				v.Build = append(v.Build, bVName)
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
	}

	s.Engine.GET("/", func(c *gin.Context) {
		var endpoints []string
		for _, route := range s.Engine.Routes() {
			if route.Path == "/" {
				continue
			}
			endpoints = append(endpoints, fmt.Sprintf("%s %s", route.Method, route.Path))
		}

		c.JSON(http.StatusOK, map[string]any{
			"hello":     "world",
			"endpoints": endpoints,
		})
	})

	/*
		TODO:
			- HTML endpoints go here
	*/

	s.Engine.Run(":8080")
}
