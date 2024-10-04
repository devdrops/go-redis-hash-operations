# Go: Redis Hash Operations

Go examples of Redis Hash operations:

- [x] HDEL
- [x] HEXISTS
- [ ] HEXPIRE
- [ ] HEXPIREAT
- [ ] HEXPIRETIME
- [x] HGET
- [x] HGETALL
- [ ] HINCRBY
- [ ] HINCRBYFLOAT
- [ ] HKEYS
- [ ] HLEN
- [x] HMGET
- [ ] HMSET
- [ ] HPERSIST
- [ ] HPEXPIRE
- [ ] HPEXPIREAT
- [ ] HPEXPIRETIME
- [ ] HPTTL
- [ ] HRANDFIELD
- [ ] HSCAN
- [ ] HSET
- [ ] HSETNX
- [ ] HSTRLEN
- [ ] HTTL
- [ ] HVALS

## Before Start

Make sure you have the following prerequisites installed correctly:

|Tool|Version\*|
|:---|:---|
|Docker|`27.1.1`|
|GNU Make|`GNU Make 4.3`|

\*: Suggested. Other versions may not work as expected.

## How to Run?

1. Clone this repository.
2. Using your terminal, go to the project's root folder.
3. In your terminal, execute the command below to start the Redis Server:

```sh
make start-redis
```

4. Also in your terminal, execute the command below to execute the Go code:

```sh
make run-go
```

If you want to peek into what's happening inside Redis, you can connect using the Redis CLI tool by executing the
command below:

```sh
make connect-redis
```

If you need to troubleshoot errors, you can check the Redis Server logs by executing the command below:

```sh
make logs-redis

# Press Ctrl + C to close
```

To stop the Redis Server, execute the command below in your terminal:

```sh
make stop-redis
```

## What is a Hash?

TBD

## Redis Hash Operations

TBD

## Obs

TBD

---

## References

- https://redis.io/docs/latest/develop/data-types/hashes/
- https://pkg.go.dev/github.com/redis/go-redis/v9
- https://redis.io/docs/latest/commands/?group=hash
- https://www.youtube.com/watch?v=-KdITaRkQ-U
-
