package agent

import (
	"net"

	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/request"
	"github.com/negasus/haproxy-spoe-go/stats"
	"github.com/negasus/haproxy-spoe-go/worker"
)

func New(handler func(*request.Request), logger logger.Logger, statter stats.Statter) *Agent {
	agent := &Agent{
		handler: handler,
		logger:  logger,
		statter: statter,
	}

	return agent
}

type Agent struct {
	handler func(*request.Request)
	logger  logger.Logger
	statter stats.Statter
}

func (agent *Agent) Serve(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			agent.statter.Count("accept", 1, stats.Tags{"error": err.Error()})
			return err
		}

		agent.statter.Count("accept", 1)
		go worker.Handle(conn, agent.handler, agent.logger)
	}
}
