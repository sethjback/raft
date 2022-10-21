.PHONY: proto
proto:
	cd proto && protoc -I ./ --go_out=./ --go-grpc_out=./ ./*.proto