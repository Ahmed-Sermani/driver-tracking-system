SCRIPTS_DIR = ./hack/scripts

.PHONY: setup setup-kind setup-ingress deploy-prometheus deploy-redis deploy-centrifugo setup-grafana-plugins help

.DEFAULT_GOAL := help

setup: setup-kind setup-ingress deploy-prometheus

deploy: deploy-prometheus deploy-redis deploy-postgres deploy-centrifugo build-image push-image deploy-app 

setup-kind:
	bash $(SCRIPTS_DIR)/setup-kind.sh

setup-ingress:
	bash $(SCRIPTS_DIR)/deploy-nginx-ingress-controller.sh

deploy-prometheus:
	bash $(SCRIPTS_DIR)/deploy-prometheus-operator.sh

deploy-redis:
	bash $(SCRIPTS_DIR)/deploy-redis.sh

deploy-postgres:
	bash $(SCRIPTS_DIR)/deploy-postgres.sh

deploy-centrifugo:
	bash $(SCRIPTS_DIR)/deploy-centrifugo.sh

deploy-app:
	bash $(SCRIPTS_DIR)/deploy-app.sh

setup-and-deploy: setup deploy

setup-grafana-plugins:
	pod=$$(kubectl get pods -n monitoring -l app.kubernetes.io/name=grafana | awk 'NR==2 {print $\$1}') && \
	kubectl exec -n monitoring  $$pod -c grafana -- grafana-cli plugins install redis-explorer-app

install-tools:
	echo installing tools
	@go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16.2
	@go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.16.2
	@go install -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
	@echo done


generate:
	@echo running code generation
	@go generate ./... && bash ./proxy/proxypb/post-generation-hook.sh
	@echo done

build-image:
	@ARCH=$(uname -m) && \
	docker buildx build -t monolith . --build-arg="GOARCH=${ARCH}"

push-image:
	echo "Loading Docker image into kind cluster..."
	@kind load docker-image monolith --name driver-tracking-system

forward-grafana:
	@kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring

help:
	@echo "This makefile contains commands for setting up and deploying a monolith application using kind, ingress, prometheus, redis and centrifugo."
	@echo "The following commands are available:"
	@echo "  setup             - Setup the kind cluster, the ingress controller and the prometheus operator"
	@echo "  setup-kind        - Setup the kind cluster using the script in $(SCRIPTS_DIR)"
	@echo "  setup-ingress     - Setup the ingress controller using the script in $(SCRIPTS_DIR)"
	@echo "  deploy-prometheus - Deploy the prometheus operator using the script in $(SCRIPTS_DIR)"
	@echo "  deploy-redis      - Deploy the redis cluster using the script in $(SCRIPTS_DIR)"
	@echo "  deploy-centrifugo - Deploy the centrifugo server using the script in $(SCRIPTS_DIR)"
	@echo "  setup-grafana-plugins - Install the grafana plugins for redis and centrifugo"
	@echo "  install-tools     - Install the tools for code generation"
	@echo "  generate          - Generate the code for grpc and openapi"
	@echo "  build-image       - Build the docker image for the monolith application"
	@echo "  push-image        - Push the docker image to the kind cluster"
