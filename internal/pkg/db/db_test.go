package db

import (
	"context"
	"testing"
)

func TestSeed(t *testing.T) {
	Init()
	err := GetDB(context.Background()).Exec("select 1").Error
	if err != nil {
		t.Fatal(err)
	}
	t.Log("done")
}
