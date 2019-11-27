package file

import (
	"fmt"
	"github.com/gottingen/felix"
	"github.com/gottingen/gekko/store/types"
)

type client struct {
	fs      felix.Felix
}

func NewFileClient(config types.Config) (types.Store, error) {
	c := new(client)
	c.fs = felix.NewOsVfs()
	return c, nil
}

func (c *client) Get(key string) ([]byte, error) {

	isDir, err := c.fs.IsDir(key)
	if err != nil {
		return []byte(""), err
	}

	if isDir {
		return []byte(""), fmt.Errorf("%s is a dir", key)
	}

	bs, errRead := c.fs.ReadFile(key)

	return bs, errRead
}

func (c *client) Set(key string, value []byte) error {
	return c.fs.WriteFile(key, value, 0644)
}

func (c *client) Close() error {
	return nil
}

func (c *client) Delete(key string) error {
	c.fs.Vfs.Remove(key)
	return nil
}