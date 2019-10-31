.PHONY: genproto
genproto:
	rm -rf genproto
	docker run -v $(PWD)/proto:/proto -v $(PWD)/genproto:/genproto mfycheng/protoc-gen-go
	mv genproto/github.com/kinecosystem/kin-api/genproto/* genproto
	rm -rf genproto/github.com
