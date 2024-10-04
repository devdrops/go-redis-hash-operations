SHELL := /bin/bash
# ----------------------------------------------------------------------------------------------------------------------
start-redis:
	@docker rm -f redis
	docker run -d --name redis \
		--network=host \
		-v $(PWD)/redis-conf/redis-data:/data \
		-v $(PWD)/redis-conf/redis.conf:/usr/local/etc/redis/redis.conf \
		redis:alpine3.20 \
		redis-server /usr/local/etc/redis/redis.conf --save 60 1 --loglevel warning

stop-redis:
	docker stop redis

connect-redis:
	docker exec -ti redis redis-cli

logs-redis:
	docker logs -t -f redis

run-go:
	@docker rm -f go
	@docker run -ti --name go \
		--network=host \
		-v $(PWD):/redis-hash \
		-v $(PWD)/go-cache:/go/pkg/mod/cache \
		-w /redis-hash \
		golang:1.23 \
		go run .

# ----------------------------------------------------------------------------------------------------------------------
.PHONY: start-redis stop-redis connect-redis logs-redis run-go
