package service

import (
	"define-komandy/internal/structs"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func readYaml[T structs.Config | structs.SecretConfig](filename string) (*T, error) {
	d, _ := os.Getwd()
	buf, err := os.ReadFile(d + filename)
	if err != nil {
		log.Fatal("Error while reading config")
	}
	var data T
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return nil, fmt.Errorf("error reading %q: %w", filename, err)
	}
	return &data, err
}
