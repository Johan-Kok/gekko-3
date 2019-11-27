package store

import (
	"github.com/gottingen/gekko/store/types"
)

type Store types.Store
type Config types.Config

type stores map[string]func(types.Config) (types.Store, error)

var storeMap = stores{}
