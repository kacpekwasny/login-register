package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	cmt "github.com/kacpekwasny/commontools"
)

var CONFIG_MAP map[string]string

func LoadConfig() {
	f, err := os.Open("logreg.conf.json")
	cmt.Pc(err)
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	cmt.Pc(err)
	err = json.Unmarshal(bytes, &CONFIG_MAP)
	cmt.Pc(err)
}
