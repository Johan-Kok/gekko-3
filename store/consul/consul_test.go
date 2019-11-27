package consul

import (
	"testing"
	"github.com/gottingen/gekko/store/types"
)

func TestConnect(t *testing.T) {

	conf := types.Config{}
	conf.Path = "gekko/store"
	conf.BackendNodes = []string{"127.0.0.1:8500"}
	client, err := NewConsul(conf)

	if client == nil {
		t.Fatalf("connect fail")
	}

	err = client.Set("abc", []byte("def"))

	if err != nil {
		t.Fatalf("set  fail %s", err.Error())
	}

	b, err1 :=client.Get("abc")
	if err1 != nil {
		t.Fatalf("get  fail %s", err.Error())
	}

	if string(b) != "def" {
		t.Fatalf("get value  fail %s", err.Error())
	}
}
