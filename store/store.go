package store

import (
	"fmt"
	"github.com/gottingen/gekko/store/consul"
	"github.com/gottingen/gekko/store/file"
	"github.com/gottingen/gekko/store/types"
)

type Store types.Store
type Config types.Config

type stores map[string]func(types.Config) (types.Store, error)

var storeMap = stores{
	"consul":consul.NewConsul,
	"file":file.NewFileClient,
	"mem":file.NewFileClient,
}

func GetStore(k string, config types.Config) (Store, error) {
	h, ok := storeMap[k]
	if !ok {
		return nil, fmt.Errorf("not support")
	}
	return h(config)
}


var EmptyConfig = types.Config{}