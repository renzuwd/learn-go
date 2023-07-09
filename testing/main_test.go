package main

import "testing"

func Test_Add(t *testing.T) {
	result := Add(1, 2)
	if result == 4 {
		t.Logf("ok")
	} else {
		t.Fail()
	}
}

func Test_Add2(t *testing.T) {
	result := Add(1, 2)
	if result == 3 {
		t.Logf("ok")
	} else {
		t.Fail()
	}
}

func Test_AddMuliti(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := Add(1, 2)
		if result == 4 {
			t.Logf("ok")
		} else {
			t.Fail()
		}
	})

	t.Run("faild", func(t *testing.T) {
		result := Add(1, 2)
		if result == 3 {
			t.Logf("ok")
		} else {
			t.Fail()
		}
	})
}
