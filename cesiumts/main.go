package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

const SampleFileName = "sample.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sampleFile := path.Join(os.Getenv("FILES_DIR"), SampleFileName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(sampleFile); errors.Is(err, os.ErrNotExist) {
			currentTime := time.Now()
			data := []byte(fmt.Sprintf("Hello efs! written at %s", currentTime))
			err := os.WriteFile(sampleFile, data, 0644)
			check(err)
		}

		dat, err := os.ReadFile(sampleFile)
		check(err)
		fmt.Print()

		io.WriteString(w, string(dat))
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(sampleFile); errors.Is(err, os.ErrNotExist) {
			err := os.Remove(sampleFile)
			check(err)
			io.WriteString(w, "true")
		}
		io.WriteString(w, "false")
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
