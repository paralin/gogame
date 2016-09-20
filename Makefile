protogen:
	protowrap -I $${GOPATH}/src \
		--gogo_out=$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/*.proto
	echo "// +build -js" >> .nobuildjs
	cat .nobuildjs gogame.pb.go > gogame.gogo.pb.go
	rm gogame.pb.go .nobuildjs
	protowrap -I $${GOPATH}/src \
		--go_out=$${GOPATH}/src \
		--proto_path $${GOPATH}/src \
		--print_structure \
		--only_specified_files \
		$$(pwd)/*.proto
	echo "// +build js" >> .buildjs
	cat .buildjs gogame.pb.go > gogame.pb.go.tmp
	mv gogame.pb.go.tmp gogame.pb.go

deps:
	go get -u github.com/square/goprotowrap/cmd/protowrap
