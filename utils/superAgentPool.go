package utils

import (
	"errors"
	"sync"

	"github.com/parnurzeal/gorequest"
)

type SuperAgentPool struct {
	mu     sync.Mutex
	agents chan *gorequest.SuperAgent
}

func NewSuperAgentPool(capacity int) (*SuperAgentPool, error) {
	if capacity < 1 {
		return nil, errors.New("invalid capacity")
	}

	pool := &SuperAgentPool{
		agents: make(chan *gorequest.SuperAgent, capacity),
	}

	for i := 0; i < capacity; i++ {
		agent := gorequest.New()
		if agent == nil {
			return nil, errors.New("unable to create gorequest.SuperAgent")
		}

		pool.agents <- agent
	}

	return pool, nil
}

func (p *SuperAgentPool) Get() *gorequest.SuperAgent {
	agent := <-p.agents
	return agent
}

func (p *SuperAgentPool) Put(agent *gorequest.SuperAgent) error {
	if agent == nil {
		return errors.New("input agent is nil")
	}

	p.agents <- agent
	return nil
}

func (p *SuperAgentPool) Len() int {
	return len(p.agents)
}
