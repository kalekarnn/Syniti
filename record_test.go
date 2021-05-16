package main

import (
	"testing"
)

func TestIsNullOrEmpty_Empty(t *testing.T) {

	emptyResult := isNullOrEmpty("")

	if !emptyResult {
		t.Error("record_test.go|TestIsNullOrEmpty_Empty", emptyResult)
	}
}

func TestIsNullOrEmpty_null(t *testing.T) {

	nullResult := isNullOrEmpty("null")

	if !nullResult {
		t.Error("record_test.go|TestIsNullOrEmpty_null", nullResult)
	}
}

func TestIsInvalidZip_5digit(t *testing.T) {

	zipResult := isInvalidZip("17565")

	if zipResult {
		t.Error("record_test.go|TestIsInvalidZip_5digit", zipResult)
	}
}

func TestIsInvalidZip_9digit(t *testing.T) {

	zipResult := isInvalidZip("17565-3456")

	if zipResult {
		t.Error("record_test.go|TestIsInvalidZip_9digit", zipResult)
	}
}

func TestIsInvalidZip_Invalid(t *testing.T) {

	zipResult := isInvalidZip("1756")

	if !zipResult {
		t.Error("record_test.go|TestIsInvalidZip_Invalid", zipResult)
	}
}

func TestIsInvalid_valid(t *testing.T) {

	var r = Record{"Amerah Lang", "5037 Providence Bouled", "44109", "8d323"}
	invalid := r.isInvalid()

	if invalid {
		t.Error("record_test.go|TestIsInvalid_valid", invalid)
	}
}

func TestIsInvalid_Invalid(t *testing.T) {

	var r = Record{"Amerah Lang", "", "44109", "8d323"}
	invalid := r.isInvalid()

	if !invalid {
		t.Error("record_test.go|TestIsInvalid_Invalid", invalid)
	}
}

func TestGetKey(t *testing.T) {

	var r = Record{"Amerah Lang", "5037 Providence Bouled", "44109", "8d323"}
	key := r.getKey()

	if key != "Amerah_Lang_5037_Providence_Bouled_44109" {
		t.Error("record_test.go|TestGetKey", key)
	}
}
