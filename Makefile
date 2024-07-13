CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

run:
	go run cmd/main.go

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs
