package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func basicAuth(fn http.HandlerFunc, jsonPath string) http.HandlerFunc {
	check := newCheckerFromJson(jsonPath)
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !check(user, pass) {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		fn(w, r)
	}
}

func newCheckerFromJson(jsonPath string) func(name, password string) bool {
	authMap := &map[string]string{}
	loadJson(jsonPath, authMap)
	return func(user, pass string) bool {
		for k, v := range *authMap {
			// log.Println(user, pass, k, v)
			if k == user && v == pass {
				return true
			}
		}
		return false
	}
}

func loadJson(path string, obj interface{}) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, obj)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(obj)
}
