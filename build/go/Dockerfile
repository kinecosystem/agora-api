FROM golang:1.14

# Base libraries
RUN apt-get update
RUN apt-get install -y protobuf-compiler

# Go proto dependencies
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Common base proto dependencies.
RUN go get -u github.com/envoyproxy/protoc-gen-validate
RUN git clone https://github.com/googleapis/googleapis

RUN mkdir /proto-common
RUN cp -r src/github.com/envoyproxy/protoc-gen-validate/validate /proto-common/validate
RUN cp -r googleapis/google /proto-common

COPY generate.sh generate.sh
RUN chmod +x generate.sh

ENTRYPOINT [ "./generate.sh" ]
