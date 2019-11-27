package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	pgmock "github.com/matthewljsmith/pgmock/pgserver"
)

// InjectData simple struct for the POST payload endpoint
type InjectData struct {
	Columns []string        `json:"cols"`
	Rows    [][]interface{} `json:"rows"`
}

func main() {

	// define some flags that get passed in
	var verbose = flag.Bool("verbose", false, "verbose - pass this to set the debug level to Debug instead of Info")
	flag.Parse()

	// if verbose set log verbosity
	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// new up the mocking server
	mock := pgmock.NewServer()

	// kick of the mocking instance
	log.Infof("starting pgmock -> 127.0.0.1:9999")
	go mock.ListenAndServe(fmt.Sprintf("127.0.0.1:9999"))

	// now kickoff the data loading api
	log.Infof("starting data-loader pg -> 127.0.0.1:9998")
	dl := gin.Default()
	dl.POST("/:hash", func(c *gin.Context) {
		var payload *InjectData
		err := c.MustBindWith(&payload, binding.JSON)
		if err != nil {
			c.AbortWithError(400, err)
		}
		err = mock.InjectQueryResponse(c.Param("hash"), payload.Columns, payload.Rows)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		c.Status(200)
	})
	dl.Run("127.0.0.1:9998")
}
