package main

import "testing"

func TestSuccess(t *testing.T) {
	t.Log("this is success")
}

func TestFail(t *testing.T) {
	t.Fatal("this is  fail")
}
