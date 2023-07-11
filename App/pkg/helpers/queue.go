package helpers

import (
	"fmt"
	"sync"
	"github.com/streadway/amqp"
)

var CconnPool *ConnPool

type ConnPool struct {
	pool   chan *amqp.Connection
	config string
	lock   sync.Mutex
}

func InitConnPool(poolSize int) (error , *ConnPool) {
	url := GetEnvStr("amqp.url")
	username := GetEnvStr("amqp.username")
	password := GetEnvStr("amqp.password")
	config := fmt.Sprintf("amqp://%s:%s@%s/",
		username,
		password,
		url,
	)

	pool := make(chan *amqp.Connection, poolSize)

	for i := 0; i < poolSize; i++ {
		conn, err := amqp.Dial(config)
		if err != nil {
			fmt.Print(err.Error())
			return err , nil
		}

		pool <- conn
	}

	CconnPool = &ConnPool{
		pool:   pool,
		config: config,
	}

	return nil , CconnPool
}

func GetConnPool() *ConnPool {
	return CconnPool
}

func (p *ConnPool) GetConn() (*amqp.Connection, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.pool) == 0 {
		conn, err := amqp.Dial(p.config)
		if err != nil {
			return nil, err
		}

		return conn, nil
	}

	return <-p.pool, nil
}

func (p *ConnPool) ReturnConn(conn *amqp.Connection) (error) {
	p.pool <- conn
	return nil
}