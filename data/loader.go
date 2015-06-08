package data

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"math/rand"
	"time"
)

func Load(name, locale, path string) string {
	fi, err := ioutil.ReadFile(fmt.Sprintf("locale/%s/%s.yml", path, locale))
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(fi), &m)
	if err != nil {
		panic(err)
	}

	stringValue, isString := m[name].(string)
	arrayValue, isArray := m[name].([]interface{})
	if isString {
		return stringValue
	} else if isArray {
		return randomElement(arrayValue)
	}

	return fmt.Sprintf("[Failed to load %v from %v-%v]", name, path, locale)
}

func randomElement(array []interface{}) string {
	rand.Seed(time.Now().UTC().UnixNano())
	a, _ := array[rand.Intn(len(array))].(string)
	return a
}
