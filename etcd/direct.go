package etcd

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func Set(key, value string) (*client.Response, error) {
	return KeysAPI.Set(context.Background(), key, value, nil)
}

func Get(key string) (*client.Response, error) {
	return KeysAPI.Get(context.Background(), key, nil)
}

func GetValue(key string) (string, error) {
	resp, err := KeysAPI.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

func GetString(key string) string {
	resp, err := GetValue(key)
	if err != nil {
		return ""
	}
	return resp
}

// key = /banerwai/mongo return multi node
// /banerwai/mongo/conn       localhost:27017
// banerwai/mongo/database    banerwai
func GetService(key string) (*client.Response, error) {
	return KeysAPI.Get(context.Background(), key, nil)
}
