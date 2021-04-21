package main

import "testing"

func TestLogs(t *testing.T) {
	err := WriteLogs("start")
	if err != nil {
		t.Error(err.Error())
	}
}
