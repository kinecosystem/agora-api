FROM node:14

# Base libraries
RUN apt-get update
RUN apt-get install -y protobuf-compiler

# Yarn proto dependencies
RUN yarn add grpc grpc-tools grpc_tools_node_protoc_ts

# Common base proto dependencies.
RUN git clone https://github.com/envoyproxy/protoc-gen-validate
RUN git clone https://github.com/googleapis/googleapis

RUN mkdir /proto-common
RUN cp -r protoc-gen-validate/validate /proto-common/validate
RUN cp -r googleapis/google /proto-common

COPY generate.sh generate.sh
RUN chmod +x generate.sh

ENTRYPOINT [ "./generate.sh" ]
