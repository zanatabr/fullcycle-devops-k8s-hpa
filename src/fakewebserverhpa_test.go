package main

import "testing"

func TestHttpSrv_DeveResultarBtesteB(t *testing.T) {
	got := greeting("teste")
	want := "<b>teste</b>"

	if got != want {
		t.Errorf("greeting(\"teste\") \n got: %v \n want: %v \n", got, want)
	}
}
