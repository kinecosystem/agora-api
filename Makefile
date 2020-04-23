USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

.PHONY: genproto
genproto:
	rm -rf genproto/*
	docker run -v $(PWD)/proto:/proto -v $(PWD)/genproto:/genproto --user $(USER_ID):$(GROUP_ID) mfycheng/protoc-gen-go
	mv genproto/github.com/kinecosystem/kin-api-internal/genproto/* genproto
	rm -rf genproto/github.com
