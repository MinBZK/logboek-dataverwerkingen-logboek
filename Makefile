SERVER_IMAGE := "ghcr.io/minbzk/logboek-dataverwerkingen-logboek"


.PHONY: all
all: proto server-image

.PHONY: server-image
server-image:
	docker build -t "${SERVER_IMAGE}" .

.PHONY: proto
proto: proto-lib-python proto-lib-go

LIB_PYTHON_DIR := libs/logboek-python/src
LIB_GO_DIR := libs/logboek-go

.PHONY: proto-lib-python
proto-lib-python:
	python3 -m grpc_tools.protoc \
		-I logboek/proto=proto \
		--python_out=${LIB_PYTHON_DIR} \
		--pyi_out=${LIB_PYTHON_DIR} \
		--grpc_python_out=${LIB_PYTHON_DIR}	 \
		proto/logboek/v1/logboek.proto

.PHONY: proto-lib-go
proto-lib-go:
	python3 -m grpc_tools.protoc \
		-I . \
		--go_out=${LIB_GO_DIR} \
		--go_opt=paths=import \
		--go_opt=Mproto/logboek/v1/logboek.proto=proto/logboek/v1 \
		--go-grpc_out=${LIB_GO_DIR} \
		--go-grpc_opt=paths=import \
		--go-grpc_opt=Mproto/logboek/v1/logboek.proto=proto/logboek/v1 \
		proto/logboek/v1/logboek.proto

.PHONY: venv
venv: .venv

.venv:
	python3 -m venv .venv
	.venv/bin/python3 -m pip install pip-tools

.PHONY: deps
deps: .venv
	.venv/bin/python3 -m pip install -r requirements.txt
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
