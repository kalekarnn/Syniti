package main

import (
	"regexp"
	"strings"
)

type Record struct {
	Name    string
	Address string
	Zip     string
	Id      string
}

func (r Record) isInvalid() bool {
	return isNullOrEmpty(r.Name) || isNullOrEmpty(r.Address) || isInvalidZip(r.Zip)
}

func isNullOrEmpty(input string) bool {
	return len(input) == 0 || input == "null"
}

func isInvalidZip(zip string) bool {
	match, _ := regexp.MatchString("^[0-9]{5}(?:-[0-9]{4})?$", zip)
	return !match
}

func (r Record) getKey() string {
	var key = r.Name+" "+r.Address+" "+r.Zip
	return strings.ReplaceAll(key, " ", "_")
}
