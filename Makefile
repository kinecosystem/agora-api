USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

.PHONY: generate
generate: gengo gennode

.PHONY: build-images
build-images:
	docker build build/go -t agora-api-builder-go
	docker build build/node -t agora-api-builder-node

.PHONY: gengo
gengo:
	rm -rf genproto/*
	docker run -v $(PWD)/proto:/proto -v $(PWD)/genproto:/genproto --user $(USER_ID):$(GROUP_ID) mfycheng/protoc-gen-go
	mv genproto/github.com/kinecosystem/agora-api/genproto/* genproto
	rm -rf genproto/github.com

.PHONY: gennode
gennode:
	docker run -v $(PWD)/proto:/proto -v $(PWD)/node:/genproto \
		--mount type=bind,source="$(PWD)/package.json",target=/package.json \
		--mount type=bind,source="$(PWD)/package.lock",target=/package.lock \
		-e "USER_ID=$(USER_ID)" -e "GROUP_ID=$(GROUP_ID)" \
		agora-api-builder-node
