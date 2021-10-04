package main

import (
	"gokv/core"
	"testing"
)

func TestAbstractKvServer(t *testing.T) {
	kvServer := new(core.AbstractKvServer)
	kvServer.Start()

	kvServer.Put("wei", "weiguangyue")

	val, ok := kvServer.Get("wei")
	if ok && val == "weiguangyue" {
		t.Log("Put,Get is Right")
	} else {
		t.Logf("Get val =%s", val)
		t.FailNow()
	}

	delOK := kvServer.Del("wei")
	if delOK {
		t.Log("Del is OK")
	} else {
		t.FailNow()
	}

	_, ok1 := kvServer.Get("wei")

	if !ok1 {
		t.Log("Del is OK")
	} else {
		t.FailNow()
	}

	kvServer.Stop()
}
