USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

.NOTPARALLEL:

.PHONY: generate
generate: gengo genpython gennode

.PHONY: build-images
build-images:
	docker build build/go -t agora-api-builder-go
	docker build build/node -t agora-api-builder-node
	docker build build/python -t agora-api-builder-python

.PHONY: gengo
gengo:
	rm -rf genproto/*
	docker run -v $(PWD)/proto:/proto -v $(PWD)/genproto:/genproto --user $(USER_ID):$(GROUP_ID) mfycheng/protoc-gen-go
	mv genproto/github.com/kinecosystem/agora-api/genproto/* genproto
	rm -rf genproto/github.com

.PHONY: genpython
genpython:
	rm -rf python/agoraapi/*
	docker run -v $(PWD)/proto:/proto -v $(PWD)/python/agoraapi:/genproto --user $(USER_ID):$(GROUP_ID) agora-api-builder-python

.PHONY: gennode
gennode:
	rm -rf node/*
	docker run -v $(PWD)/proto:/proto -v $(PWD)/node:/genproto \
		--mount type=bind,source="$(PWD)/package.json",target=/package.json \
		--mount type=bind,source="$(PWD)/yarn.lock",target=/yarn.lock \
		-e "USER_ID=$(USER_ID)" -e "GROUP_ID=$(GROUP_ID)" \
		agora-api-builder-node

