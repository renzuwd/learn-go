package model

import "testing"

func Test_GetUser(t *testing.T) {
	result := GetUser()
	if result == "GET USER2" {
		t.Logf("ok")
	} else {
		t.Fail()
	}
}
