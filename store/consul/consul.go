package consul

import (
	"fmt"
	"github.com/gottingen/gekko/store/types"
	"github.com/gottingen/gekko/store/util"
	"github.com/hashicorp/consul/api"
)

type consulClient struct {
	client *api.KV
	prefix string
}

func (cc *consulClient) Set(key string, value []byte) error {
	if err := util.CheckKeyAndValue(key, value); err != nil {
		return err
	}
	if cc.prefix != "" {
		key = cc.prefix + "/" + key
	}

	kvPair := api.KVPair{
		Key:   key,
		Value: value,
	}

	_, err := cc.client.Put(&kvPair, nil)
	if err != nil {
		return err
	}

	return nil
}

func (cc *consulClient) Get(key string) ([]byte, error) {
	if err := util.CheckKey(key); err != nil {
		return []byte(""), err
	}

	if cc.prefix != "" {
		key = cc.prefix + "/" + key
	}
	kvPair, _, err := cc.client.Get(key, nil)
	if err != nil {
		return []byte(""), err
	}
	// If no value was found return false
	if kvPair == nil {
		return []byte(""), fmt.Errorf("no value")
	}

	return kvPair.Value, nil
}

func (cc *consulClient) Delete(key string) error {
	if err := util.CheckKey(key); err != nil {
		return err
	}

	if cc.prefix != "" {
		key = cc.prefix + "/" + key
	}
	_, err := cc.client.Delete(key, nil)
	return err
}

func (cc *consulClient) Close() error {
	return nil
}

func NewConsul(config types.Config) (types.Store, error) {
	conf := api.DefaultConfig()

	conf.Scheme = config.Scheme

	if len(config.BackendNodes) > 0 {
		conf.Address = config.BackendNodes[0]
	}

	if config.BasicAuth {
		conf.HttpAuth = &api.HttpBasicAuth{
			Username: config.Username,
			Password: config.Password,
		}
	}

	if config.ClientCert != "" && config.ClientKey != "" {
		conf.TLSConfig.CertFile =  config.ClientCert
		conf.TLSConfig.KeyFile =  config.ClientKey
	}
	if config.ClientCaKeys != "" {
		conf.TLSConfig.CAFile = config.ClientCaKeys
	}

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &consulClient{client.KV(), config.Path}, nil
}
