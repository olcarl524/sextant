package sextant_test

import (
	"os"
	"path"
	"sextant"
	"testing"
)

var SextantContext *sextant.Definition

func TestSextantSet(t *testing.T) {
	tset := sextant.New().WithDatabase(path.Join("tmp", "writeTest"))
	err := tset.Set("key", "value")
	if err != nil {
		t.Errorf("Got error from upstream method<Set> err:%s", err)
	}
	if tget, errG := tset.Get("key"); tget != "value" {
		if errG != nil {
			t.Errorf("Got error from upstream method<Get> err:%s", errG)
		}
		t.Errorf("Was expecting `value` as value but instead got: %s", tget)
	}
}

func TestSextantGet(t *testing.T) {
	v, err := SextantContext.Get("read")
	if v != "working" {
		if err != nil {
			t.Errorf("Got error from upstream method<Get> err:%s", err)
		}
		t.Errorf("Was expecting `working` as value but instead got: %s", v)
	}
}

func TestMain(m *testing.M) {
	setup()
	runnable := m.Run()
	os.Exit(runnable)
}

func setup() {
	SextantContext = sextant.New().WithDatabase(path.Join("tmp", "readTest"))
	SextantContext.Set("read", "working")
}
