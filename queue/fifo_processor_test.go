package queue_test

import (
	"testing"

	"github.com/WatchBeam/redutil/conn"
	"github.com/WatchBeam/redutil/queue"
	"github.com/WatchBeam/redutil/test"
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/suite"
)

type FIFOProcessorTest struct {
	*test.RedisSuite
}

func TestFIFOProcessorSuite(t *testing.T) {
	pool, _ := conn.New(conn.ConnectionParam{
		Address: "127.0.0.1:6379",
	}, 1)

	suite.Run(t, &FIFOProcessorTest{test.NewSuite(pool)})
}

func (suite *FIFOProcessorTest) assertOrder(cnx redis.Conn) {
	first, e1 := queue.FIFO.Pull(cnx, "keyspace")
	second, e2 := queue.FIFO.Pull(cnx, "keyspace")
	third, e3 := queue.FIFO.Pull(cnx, "keyspace")

	suite.Assert().Equal([]byte("first"), first)
	suite.Assert().Equal([]byte("second"), second)
	suite.Assert().Equal([]byte("third"), third)

	suite.Assert().Nil(e1)
	suite.Assert().Nil(e2)
	suite.Assert().Nil(e3)
}

func (suite *FIFOProcessorTest) TestPullToOrder() {
	cnx := suite.Pool.Get()
	defer cnx.Close()

	queue.FIFO.Push(cnx, "keyspace", []byte("first"))
	queue.FIFO.Push(cnx, "keyspace", []byte("second"))
	queue.FIFO.Push(cnx, "keyspace2", []byte("third"))

	queue.FIFO.PullTo(cnx, "keyspace2", "keyspace")

	suite.assertOrder(cnx)
}

func (suite *FIFOProcessorTest) TestProcessingOrder() {
	cnx := suite.Pool.Get()
	defer cnx.Close()

	queue.FIFO.Push(cnx, "keyspace", []byte("first"))
	queue.FIFO.Push(cnx, "keyspace", []byte("second"))
	queue.FIFO.Push(cnx, "keyspace", []byte("third"))

	suite.assertOrder(cnx)
}

func (suite *FIFOProcessorTest) TestConcats() {
	cnx := suite.Pool.Get()
	defer cnx.Close()

	queue.FIFO.Push(cnx, "keyspace", []byte("first"))
	queue.FIFO.Push(cnx, "keyspace2", []byte("second"))
	queue.FIFO.Push(cnx, "keyspace2", []byte("third"))

	suite.Assert().Nil(queue.FIFO.Concat(cnx, "keyspace2", "keyspace"))
	suite.Assert().Nil(queue.FIFO.Concat(cnx, "keyspace2", "keyspace"))
	suite.Assert().Equal(redis.ErrNil, queue.FIFO.Concat(cnx, "keyspace2", "keyspace"))

	suite.assertOrder(cnx)
}