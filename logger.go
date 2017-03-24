package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"
)

var (
	logger *log.Logger
	logDir = flag.String("logdir", "/tmp/plane", "Directory in which to log requests")
)

func init() {
	// Initialize the logger
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// Initialize the directories
	// First the request directory
	reqDirPath := filepath.Join(*logDir, "req")
	err := os.MkdirAll(reqDirPath, 0755)
	if err != nil {
		log.Printf("init: couldn't create directory %s error: %v", reqDirPath, err)
	}
	// Now the request directory
	respDirPath := filepath.Join(*logDir, "resp")
	err = os.MkdirAll(respDirPath, 0755)
	if err != nil {
		log.Printf("init: couldn't create directory %s error: %v", respDirPath, err)
	}
}

const logHttpTimeFormat string = "2006-01-02T150405.999999999"

func logRequest(req *http.Request) error {
	// Dump the request
	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		return fmt.Errorf("logRequest: DumpRequest failed: %v", err)
	}

	// Open a file
	fileName := time.Now().Format(logHttpTimeFormat)
	filePath := filepath.Join(*logDir, "req", fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("logRequest: opening file failed: %v", err)
	}
	defer f.Close()

	// Copy
	buf := bytes.NewReader(data)
	_, err = io.Copy(f, buf)
	if err != nil {
		return fmt.Errorf("logRequest: writing to file failed: %v", err)
	}

	return nil
}

/*
func logResponse(resp *http.Response) error {
	// Dump the request
	data,err := httputil.DumpResponse(resp, true)
	if err != nil{
		return fmt.Errorf("logResponse: DumpResponse failed: %v",err)
	}

	// Open a file
	fileName := time.Now().Format(logHttpTimeFormat)
	filePath := filepath.Join(*logDir,"resp",fileName)
	f,err := os.Create(filePath)
	if err != nil{
		return fmt.Errorf("logResponse: opening file failed: %v",err)
	}
	defer f.Close()

	// Copy
	buf := bytes.NewReader(data)
	_,err = io.Copy(f,buf)
	if err != nil{
		return fmt.Errorf("logResponse: writing to file failed: %v",err)
	}

	return nil
}*/
