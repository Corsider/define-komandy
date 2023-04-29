package service

import (
	"define-komandy/internal/structs"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)

func First[T, U any](val T, _ U) T {
	return val
}

func ReadYaml[T structs.Config | structs.SecretConfig](filename string) (*T, error) {
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

func StringToArray(s string) []int {
	res := strings.Split(s, ",")
	resint := []int{}
	for _, elem := range res {
		resint = append(resint, First(strconv.Atoi(elem)))
	}
	return resint
}

func ArrayToString(a []int) string {
	res := ""
	for _, elem := range a {
		res += strconv.Itoa(elem)
	}
	return res
}
