all: lint demo

.PHONY: devoir10
devoir10: cmd
	rm devoir10
	go build -o devoir10 .

lint:
	gofmt -d -e -s .
	find . -type f -name "*.go" -exec goimports -d {} \;
	golint -min_confidence=0.3 ./...
	staticcheck ./...
	golangci-lint run

.PHONY: demo
demo: devoir10
	reset
	rm -fr 10.db
	# Beginning
	./devoir10 migrate up
	./devoir10 migrate version
	# V1 operations
	./devoir10 v1 load
	./devoir10 v1 lscompany
	./devoir10 v1 rmcompany 3
	./devoir10 v1 chcompany 1 2; true # The "; true" forces success on failing commands
	./devoir10 v1 chcompany 3 1; true
	./devoir10 v1 chcompany 1 3

	# Migration to V2
	./devoir10 migrate up
	./devoir10 v2 lscompany
	./devoir10 v2 rmperson 31; true
	./devoir10 v2 rmperson 21
	./devoir10 v2 -v chperson 11 John Doe
	./devoir10 v2 chperson 21 Max Doe; true
	./devoir10 v2 lscompany

	# Migration back to V1
	./devoir10 migrate down
	./devoir10 v1 lscompany
