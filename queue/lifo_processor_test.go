package queue_test

import (
	"testing"

	"github.com/WatchBeam/redutil/conn"
	"github.com/WatchBeam/redutil/queue"
	"github.com/WatchBeam/redutil/test"
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/suite"
)

type LIFOProcessorTest struct {
	*test.RedisSuite
}

func TestLIFOProcessorSuite(t *testing.T) {
	pool, _ := conn.New(conn.ConnectionParam{
		Address: "127.0.0.1:6379",
	}, 1)

	suite.Run(t, &LIFOProcessorTest{test.NewSuite(pool)})
}

func (suite *LIFOProcessorTest) assertOrder(cnx redis.Conn) {
	first, _ := queue.LIFO.Pull(cnx, "keyspace")
	second, _ := queue.LIFO.Pull(cnx, "keyspace")
	third, _ := queue.LIFO.Pull(cnx, "keyspace")
	suite.Assert().Equal([]byte("third"), first)
	suite.Assert().Equal([]byte("second"), second)
	suite.Assert().Equal([]byte("first"), third)
}

func (suite *LIFOProcessorTest) TestProcessingOrder() {
	cnx := suite.Pool.Get()
	defer cnx.Close()

	queue.LIFO.Push(cnx, "keyspace", []byte("first"))
	queue.LIFO.Push(cnx, "keyspace", []byte("second"))
	queue.LIFO.Push(cnx, "keyspace", []byte("third"))

	suite.assertOrder(cnx)
}

func (suite *LIFOProcessorTest) TestConcats() {
	cnx := suite.Pool.Get()
	defer cnx.Close()

	queue.LIFO.Push(cnx, "keyspace", []byte("first"))
	queue.LIFO.Push(cnx, "keyspace2", []byte("second"))
	queue.LIFO.Push(cnx, "keyspace2", []byte("third"))

	suite.Assert().Nil(queue.LIFO.Concat(cnx, "keyspace2", "keyspace"))
	suite.Assert().Nil(queue.LIFO.Concat(cnx, "keyspace2", "keyspace"))
	suite.Assert().Equal(redis.ErrNil, queue.LIFO.Concat(cnx, "keyspace2", "keyspace"))

	suite.assertOrder(cnx)
}
