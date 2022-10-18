package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

const SampleFile = "/mnt/lambda/file"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(SampleFile); errors.Is(err, os.ErrNotExist) {
			currentTime := time.Now()
			data := []byte(fmt.Sprintf("Hello efs! written at %s", currentTime))
			err := os.WriteFile(SampleFile, data, 0644)
			check(err)
		}

		dat, err := os.ReadFile(SampleFile)
		check(err)
		fmt.Print()

		io.WriteString(w, string(dat))
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(SampleFile); errors.Is(err, os.ErrNotExist) {
			err := os.Remove(SampleFile)
			check(err)
			io.WriteString(w, "true")
		}
		io.WriteString(w, "false")
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
