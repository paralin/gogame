protogen:
	protowrap -I $${GOPATH}/src \
		--gogo_out=plugins=grpc:$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/*.proto

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap
