#!/bin/bash -e
FILES="/proto-common/validate/validate.proto"
for f in $(find /proto/ -type f -name "*.proto"); do
    FILES="$FILES $f"
done

yarn add @grpc/grpc-js grpc-tools grpc_tools_node_protoc_ts

PROTO_DEST=../genproto

# JavaScript code generation
yarn run grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:${PROTO_DEST} \
    --grpc_out=${PROTO_DEST} \
    --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
    -I /proto-common:/proto \
    $FILES

# TypeScript code generation
yarn run grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=${PROTO_DEST} \
    -I /proto-common:/proto \
    $FILES

chown -R $USER_ID:$GROUP_ID ../genproto
