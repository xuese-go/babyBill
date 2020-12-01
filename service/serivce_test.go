package service

import "testing"

func TestSave(t *testing.T) {
	if err := Save("a", "b", "c"); err != nil {
		t.Log(err.Error())
	}
}
