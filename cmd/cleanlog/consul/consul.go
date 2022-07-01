package consul

import (
	consul "github.com/hashicorp/consul/api"
)

func ConsulClient() *consul.Client {
	client, err := consul.NewClient(&consul.Config{
		Address:    "consul.tyc.local",
		Scheme:     "http",
		Datacenter: "",
		Transport:  nil,
		HttpClient: nil,
		HttpAuth:   nil,
		WaitTime:   0,
		Token:      "",
		TokenFile:  "",
		Namespace:  "",
		TLSConfig:  consul.TLSConfig{},
	})
	if err != nil {
		panic(err)
	}
	return client
}

func ConsulPutKV(key string, value string) error {
	client := ConsulClient()
	kv := client.KV()
	p := &consul.KVPair{
		Key:         key,
		CreateIndex: 0,
		ModifyIndex: 0,
		LockIndex:   0,
		Flags:       0,
		Value:       []byte(value),
	}
	_, err := kv.Put(p, nil)
	return err
}

func ConsulGetKV(k string) (*consul.KVPair, *consul.QueryMeta, error) {
	client := ConsulClient()
	kv := client.KV()
	return kv.Get(k, nil)

}
func ConsulDelKV(key string) error {
	client := ConsulClient()
	kv := client.KV()

	_, err := kv.Delete(key, nil)

	return err
}

func ConsulListKey(key string) (consul.KVPairs, *consul.QueryMeta, error) {
	client := ConsulClient()
	kv := client.KV()
	return kv.List(key, nil)
}
