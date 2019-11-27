package file

import (
	"github.com/gottingen/gekko/store/types"
	"testing"
)

func TestConnect(t *testing.T) {
	config := types.Config{}
	client, err := NewFileClient(config)

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

