package auto_render

import (
	"container/list"
	"sync"
)

type Server struct {
	addr     string
	priority int
}

type ServerPool struct {
	sync.Mutex
	started     bool
	servers     map[string]*Server
	idleServers *list.List
	busyServers *list.List
}

func NewServerPool() *ServerPool {
	m := make(map[string]*Server)
	bl := list.New()
	il := list.New()
	return &ServerPool{servers: m, idleServers: il, busyServers: bl}
}

func (this *ServerPool) Start() {
	this.started = true
}

func (this *ServerPool) Add(s *Server) bool {
	this.Lock()
	defer this.Unlock()
	this.servers[s.addr] = s
	this.idleServers.PushFront(s)
}

func (this *ServerPool) Remove(addr string) bool {
	this.Lock()
	defer this.Unlock()
	for _, server := range this.servers {
		if server.addr == addr {
			break
		}
	}
}
