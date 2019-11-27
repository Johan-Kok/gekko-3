package file

import (
	"testing"
)

func TestConnect(t *testing.T) {

	client, err := NewFileClient()

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

