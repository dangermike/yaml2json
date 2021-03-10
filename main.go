package main

import (
	"io"
	"os"

	"log"

	jsoniter "github.com/json-iterator/go"
	// "gopkg.in/yaml.v2"
	"github.com/goccy/go-yaml"
)

func getSource() (io.Reader, error) {
	if len(os.Args) > 1 && os.Args[1] != "-" {
		return os.Open(os.Args[1])
	}
	return os.Stdin, nil
}

func main() {
	if err := doWork(); err != nil {
		log.Fatalf("yaml2json failed: %s", err.Error())
		os.Exit(1)
	}
}

func doWork() error {
	r, err := getSource()
	if err != nil {
		return err
	}

	dec := yaml.NewDecoder(r)
	for {
		data := map[string]interface{}{}

		if err := dec.Decode(&data); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		enc := jsoniter.NewEncoder(os.Stdout)
		if err := enc.Encode(data); err != nil {
			return err
		}
	}
}
