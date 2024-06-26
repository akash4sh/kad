PREFIX := kad
SERVER_APP_NAME := server
AGENT_APP_NAME := agent
DEPLOYMENT_WORKER_APP_NAME := deployment-worker
CONFIG_WORKER_APP_NAME := config-worker
BUILD := 0.1.1

.PHONY: gen-protoc
gen-protoc:
	mkdir -p server/pkg/pb/serverpb
	mkdir -p capten/common-pkg/pb/captenpluginspb
	mkdir -p server/pkg/pb/agentpb
	mkdir -p server/pkg/pb/captenpluginspb
	mkdir -p server/pkg/pb/pluginstorepb
	mkdir -p capten/common-pkg/pb/vaultcredpb
	mkdir -p capten/common-pkg/pb/agentpb
	mkdir -p capten/common-pkg/pb/clusterpluginspb
	mkdir -p server/pkg/pb/clusterpluginspb
	mkdir -p capten/common-pkg/pb/pluginstorepb

	cd proto && protoc --go_out=../server/pkg/pb/serverpb/ --go_opt=paths=source_relative \
    		--go-grpc_out=../server/pkg/pb/serverpb  --go-grpc_opt=paths=source_relative \
    		./server.proto

	cd proto && protoc --go_out=../server/pkg/pb/agentpb --go_opt=paths=source_relative \
    		--go-grpc_out=../server/pkg/pb/agentpb --go-grpc_opt=paths=source_relative \
    		./agent.proto
	
	cd proto && protoc --go_out=../capten/common-pkg/pb/agentpb --go_opt=paths=source_relative \
    		--go-grpc_out=../capten/common-pkg/pb/agentpb --go-grpc_opt=paths=source_relative \
    		./agent.proto

	cd proto && protoc --go_out=../server/pkg/pb/captenpluginspb --go_opt=paths=source_relative \
    		--go-grpc_out=../server/pkg/pb/captenpluginspb --go-grpc_opt=paths=source_relative \
    		./capten_plugins.proto

	cd proto && protoc --go_out=../capten/common-pkg/pb/captenpluginspb --go_opt=paths=source_relative \
    		--go-grpc_out=../capten/common-pkg/pb/captenpluginspb --go-grpc_opt=paths=source_relative \
    		./capten_plugins.proto

	cd proto && protoc --go_out=../capten/common-pkg/pb/vaultcredpb --go_opt=paths=source_relative \
    		--go-grpc_out=../capten/common-pkg/pb/vaultcredpb --go-grpc_opt=paths=source_relative \
    		./vault_cred.proto

	cd proto && protoc --go_out=../server/pkg/pb/pluginstorepb --go_opt=paths=source_relative \
    		--go-grpc_out=../server/pkg/pb/pluginstorepb --go-grpc_opt=paths=source_relative \
    		./plugin_store.proto

	cd proto && protoc --go_out=../capten/common-pkg/pb/pluginstorepb --go_opt=paths=source_relative \
    		--go-grpc_out=../capten/common-pkg/pb/pluginstorepb --go-grpc_opt=paths=source_relative \
    		./plugin_store.proto
	cd proto && protoc --go_out=../capten/common-pkg/pb/clusterpluginspb --go_opt=paths=source_relative \
    		--go-grpc_out=../capten/common-pkg/pb/clusterpluginspb --go-grpc_opt=paths=source_relative \
    		./cluster_plugins.proto
	cd proto && protoc --go_out=../server/pkg/pb/clusterpluginspb --go_opt=paths=source_relative \
    		--go-grpc_out=../server/pkg/pb/clusterpluginspb --go-grpc_opt=paths=source_relative \
    		./cluster_plugins.proto

docker-build-server:
	# The prefix for server to changed either as server or intelops-kad-server
	docker build --platform=linux/amd64 -f dockerfiles/server/Dockerfile -t ${PREFIX}-${SERVER_APP_NAME}:${BUILD} .

docker-build-kad: docker-build-agent docker-build-deployment docker-build-config

docker-build-agent:
	docker build -f dockerfiles/agent/Dockerfile -t ${PREFIX}-${AGENT_APP_NAME}:${BUILD} .

docker-build-deployment:
	docker build -f dockerfiles/deployment-worker/Dockerfile -t ${PREFIX}-${DEPLOYMENT_WORKER_APP_NAME}:${BUILD} .

docker-build-config:
	docker build -f dockerfiles/config-worker/Dockerfile -t ${PREFIX}-${CONFIG_WORKER_APP_NAME}:${BUILD} .

docker-build: docker-build-kad docker-build-server
