PROTO_DIR = ./proto/defs/*

protos: $(PROTO_DIR)
	for dir in $^ ; do protoc \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I ${GOPATH}/src/github.com/golang/protobuf/ptypes/struct \
		-I $${dir} \
		--proto_path=. \
		--go_out=plugins=grpc,paths=source_relative:./proto/lang/go/$$(echo $${dir} | sed 's/proto\/defs//') \
		--csharp_out=./proto/lang/csharp/$$(echo $${dir} | sed 's/proto\/defs//') \
		--python_out=./proto/lang/python/$$(echo $${dir} | sed 's/proto\/defs//') \
		--cpp_out=./proto/lang/cpp/$$(echo $${dir} | sed 's/proto\/defs//') \
		--js_out=./proto/lang/javascript/$$(echo $${dir} | sed 's/proto\/defs//') \
		$${dir}/*.proto ; done