package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/charmbracelet/log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:embed templates/*
var templateFS embed.FS

const (
	v1 = iota + 1
)

type BumpyServer struct {
	Engine *gin.Engine
}

func New() BumpyServer {
	router := gin.Default()
	router.SetHTMLTemplate(
		mustParseTemplates(templateFS, "templates/*"),
	)

	return BumpyServer{
		Engine: router,
	}
}

func (s BumpyServer) Run() {
	apiV1Path := fmt.Sprintf("/api/v%d", v1)
	apiV1 := s.Engine.Group(apiV1Path)
	{
		producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
		if err != nil {
			log.Errorf("%s", err)
			return
		}

		apiV1.POST("/major/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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

			payload := map[string]any{
				"bump":         "major",
				"version":      v.String(),
				"package_name": packageName,
			}

			jsonBytes, err := json.Marshal(payload)
			if err != nil {
				log.Errorf("%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serialize response"})
				return
			}

			topic := "bumpy_send"

			err = producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic},
				Key:            []byte(uuid.New().String()),
				Value:          jsonBytes,
			}, nil)

			if err != nil {
				log.Errorf("%s", err.Error())
				return
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.GET("/major/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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

			payload := map[string]any{
				"bump":         "major",
				"version":      v.String(),
				"package_name": packageName,
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.POST("/minor/:version", func(c *gin.Context) {

			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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

			payload := map[string]string{
				"bump":         "minor",
				"version":      v.String(),
				"package_name": packageName,
			}

			jsonBytes, err := json.Marshal(payload)
			if err != nil {
				log.Errorf("%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serialize response"})
				return
			}

			topic := "bumpy_send"

			err = producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic},
				Key:            []byte(uuid.New().String()),
				Value:          jsonBytes,
			}, nil)

			if err != nil {
				log.Errorf("%s", err.Error())
				return
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.GET("/minor/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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

			payload := map[string]string{
				"bump":         "minor",
				"version":      v.String(),
				"package_name": packageName,
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.POST("/patch/:version", func(c *gin.Context) {

			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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
				log.Errorf("%s", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			payload := map[string]string{
				"bump":         "patch",
				"version":      v.String(),
				"package_name": packageName,
			}

			jsonBytes, err := json.Marshal(payload)
			if err != nil {
				log.Errorf("%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serialize response"})
				return
			}

			topic := "bumpy_send"

			err = producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic},
				Key:            []byte(uuid.New().String()),
				Value:          jsonBytes,
			}, nil)

			if err != nil {
				log.Errorf("%s", err.Error())
				return
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.GET("/patch/:version", func(c *gin.Context) {
			inputVersion := c.Param("version")

			v, err := semver.Make(inputVersion)
			if err != nil {
				log.Errorf("%s", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var packageName string

			if len(c.Query("package_name")) != 0 {
				_, err := url.Parse(c.Query("package_name"))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				packageName = c.Query("package_name")
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
				log.Errorf("%s", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			payload := map[string]string{
				"bump":         "patch",
				"version":      v.String(),
				"package_name": packageName,
			}

			c.JSON(http.StatusOK, payload)
		})

		apiV1.GET("/endpoints", func(c *gin.Context) {
			var endpoints []string

			setEndpoints(&endpoints, s.Engine.Routes())

			c.JSON(http.StatusOK, map[string]any{
				"endpoints": endpoints,
			})
		})
	}

	s.Engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]string{})
	})

	s.Engine.Run(":8080")
}

func setEndpoints(e *[]string, routeInfo gin.RoutesInfo) {
	for _, route := range routeInfo {
		if route.Path == "/" {
			continue
		}
		*e = append(*e, fmt.Sprintf("%s %s", route.Method, route.Path))
	}
}

func mustParseTemplates(fs embed.FS, pattern string) *template.Template {
	tmpl := template.New("")
	entries, err := fs.ReadDir(strings.TrimSuffix(pattern, "/*"))
	if err != nil {
		log.Fatal("failed to read embedded template dir", "error", err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		content, err := fs.ReadFile(strings.TrimSuffix(pattern, "/*") + "/" + e.Name())
		if err != nil {
			log.Fatal("failed to read embedded template", "file", e.Name(), "error", err)
		}
		tmpl, err = tmpl.New(e.Name()).Parse(string(content))
		if err != nil {
			log.Fatal("failed to parse embedded template", "file", e.Name(), "error", err)
		}
	}
	return tmpl
}
