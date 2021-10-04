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

type AbstractKvServer struct {
	kv map[string]string
}

func (server *AbstractKvServer) Start() {
	log.Println("start")
	server.kv = make(map[string]string)
}

func (server *AbstractKvServer) Stop() {
	log.Println("stop")
}

func (server *AbstractKvServer) Put(name string, val string) bool {
	server.kv[name] = val
	return true
}

func (server *AbstractKvServer) Del(name string) bool {
	delete(server.kv, name)
	return true
}

func (server *AbstractKvServer) Get(name string) (string, bool) {
	val, ok := server.kv[name]
	return val, ok
}
