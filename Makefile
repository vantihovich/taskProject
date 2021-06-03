#build information
LDFLAGS       := -w -s
BINDIR        := $(CURDIR)/bin
OS            ?= windows
PROTODIR      := $(CURDIR)/proto

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
	
protogen:	
	protoc --go_out=proto --proto_path=$(PROTODIR) --go_opt=paths=source_relative \
    --go-grpc_out=proto --proto_path=$(PROTODIR) --go-grpc_opt=paths=source_relative \
    $(PROTODIR)/*.proto
    