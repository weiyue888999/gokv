package core

import (
	"log"
)

type KvServer interface {
	/**
	** 启动
	**/
	Start()

	/**
	** 停止
	**/
	Stop()

	/**
	** 放入key,val
	**/
	Put(name string, val string) bool

	/**
	** 删除key
	**/
	Del(name string) bool

	/**
	** 获取key的val
	**/
	Get(name string) (string, bool)
}

type DefaultKvServer struct {
	kv map[string]string
}

func (server *DefaultKvServer) Start() {
	log.Println("start")
	server.kv = make(map[string]string)
}

func (server *DefaultKvServer) Stop() {
	log.Println("stop")
}

func (server *DefaultKvServer) Put(name string, val string) bool {
	server.kv[name] = val
	return true
}

func (server *DefaultKvServer) Del(name string) bool {
	delete(server.kv, name)
	return true
}

func (server *DefaultKvServer) Get(name string) (string, bool) {
	val, ok := server.kv[name]
	return val, ok
}
