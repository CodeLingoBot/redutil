# Changelog

## pubsub

### 2.4.2 (02-19-2016)
 * Extract a `worker.Worker` interface, and create a `worker.DefaultWorker`
   implementation.

### 2.4.1 (02-17-2016)
 * Fix a critical issue in the worker package which prevented janitors from
   being able to clean up dead workers.

### 2.4 (02-14-2016)
 * Implement the `worker` package.

### 2.3.1 (02-13-2016)
 * **Breaking Change**: Allow the specification of a timeout parameter to
   several (previously) infinitely-blocking methods.

### 2.3 (01-30-2016)
 * Implement the `queue` package.

### 2.2 (01-26-2016)
 * Implement the `heartbeat` package.

### 2.1 (01-23-2016)

 * **Breaking Change**: `ConnectionParam` has moved from the `pubsub` package to
   the `conn` package.
 * **Breaking Change**: `pubsub.New` no longer takes a `ConnectionParam`, rather
   it takes a `*redis.Pool` and a `conn.ReconnectPolicy`.

### 2.0 (25-08-2015) rc

 * **Breaking Change**: New() now takes a ConnectionParam value rather than a pointer.
 * **Breaking Change**: GetState() now returns a uint8 rather than a user-defined type, for greater compatibility with [fsm](https://github.com/mixer/fsm). <small>_mumble mumble generics_</small>
 * Fix potential data races on the internal subscription registry.
 * Fix potential data race resulting in subscription duplication during multiple reconnections.
 * Allow specification of connection timeout (deadlines).
 * Allow specification of reconnection policies.
 * Cause subscription, unsubscriptions, and teardowns to happen more quickly.
 * Improve events system for increased flexibility.
 * Significantly improve conciseness and speed.


### 1.1 (25-08-2015)

 * **Breaking Change**: New() now takes a *ConnectionParam struct as its first argument.
 * Add password authentication options (from @janeczku).
 * Fix failing tests in Go 1.3
 * Prevent paniking when tearing down a client which was not set up.


### 1.0 (07-04-2015)

Initial
