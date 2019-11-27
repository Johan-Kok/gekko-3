package mem

import (
	"github.com/gottingen/gekko/store/types"
	"github.com/gottingen/gekko/store/util"
	"sync"
)

type memClient struct {
	mem     map[string][]byte
	lock  	sync.RWMutex
}

func NewMemClient() (types.Store, error) {
	m := new(memClient)
	m.mem = make(map[string][]byte)
	return m, nil
}

func (m *memClient) Set(key string, value []byte) error {
	if err := util.CheckKeyAndValue(key, value); err != nil {
		return err
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.mem[key] = value
	return nil
}

func (m *memClient) Get(key string) ([]byte, error)   {
	if err := util.CheckKey(key); err != nil {
		return []byte(""), err
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.mem[key]
	if ok {
		return v, nil
	}
	return []byte(""), nil
}


func (m *memClient) Delete(key string) error {
	if err := util.CheckKey(key); err != nil {
		return err
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.mem, key)
	return nil
}

func (m *memClient) Close() error {
	return nil
}