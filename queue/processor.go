package queue

import "github.com/garyburd/redigo/redis"

// Processor is an interface to a type encapsulating the interaction between a
// queue.ByteQueue and a datastructure in Redis.
type Processor interface {
	// Push pushes a given `payload` into the keyspace at `key` over the
	// given `redis.Conn`. This function should block until the item can
	// succesfully be confirmed to have been pushed.
	Push(conn redis.Conn, key string, payload []byte) (err error)

	// Pull pulls a given `payload` from the keyspace at `key` over the
	// given `redis.Conn`. This function should block until a timeout has
	// elapsed, or an item is available.
	Pull(conn redis.Conn, key string) (payload []byte, err error)

	// Moves all elements from the src queue to the end of the destination
	// It should return a redis.ErrNil when the source queue is empty.
	Concat(conn redis.Conn, src, dest string) (err error)
}
