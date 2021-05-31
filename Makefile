#build information
LDFLAGS       := -w -s
BINDIR        := $(CURDIR)/bin
OS            ?= windows
PROTODIR      := $(CURDIR)/proto2

#service and environment information
BASE_NAME     ?= auth
SERVICE_NAME  ?= authorisation
#ENVIRONMENT   ?= dev

build:
	$(info Building binary to bin/$(SERVICE_NAME))
	@CGO_ENABLED=0  go build -o $(BINDIR)/$(SERVICE_NAME) -installsuffix cgo -ldflags '$(LDFLAGS)' ./cmd

run: build
	$(info Running LOG_LEVEL=$(LOG_LEVEL) APP_PORT=$(APP_PORT) $(BINDIR)/$(SERVICE_NAME))
	@LOG_LEVEL=$(LOG_LEVEL) APP_PORT=$(APP_PORT) $(BINDIR)/$(SERVICE_NAME)
	
proto2:
	$(info Building pb.go and _grpc.pb.go to proto2/)
	protoc $(PROTODIR)/creds2.proto --go_out=. --go_opt=paths=source_relative 
	--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false