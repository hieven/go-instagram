package utils

import (
	"errors"
	"sync"

	"github.com/parnurzeal/gorequest"
)

type SuperAgentPool struct {
	Capacity int
	Mu       sync.Mutex
	Agents   chan *gorequest.SuperAgent
}

func NewSuperAgentPool(capacity int) (*SuperAgentPool, error) {
	if capacity < 1 {
		return nil, errors.New("invalid capacity")
	}

	pool := &SuperAgentPool{
		Capacity: capacity,
		Agents:   make(chan *gorequest.SuperAgent, capacity),
	}

	for i := 0; i < pool.Capacity; i++ {
		agent := gorequest.New()
		if agent == nil {
			return nil, errors.New("unable to create gorequest.SuperAgent")
		}

		pool.Agents <- agent
	}

	return pool, nil
}

func (p *SuperAgentPool) Get() *gorequest.SuperAgent {
	agent := <-p.Agents
	return agent
}

func (p *SuperAgentPool) Put(agent *gorequest.SuperAgent) error {
	if agent == nil {
		return errors.New("input agent is nil")
	}

	p.Agents <- agent
	return nil
}

func (p *SuperAgentPool) Len() int {
	return len(p.Agents)
}
