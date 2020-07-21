#!/bin/bash -e
FILES="/proto-common/validate/validate.proto"
for f in $(find /proto/ -type f -name "*.proto"); do
    FILES="$FILES $f"
done

for f in $FILES; do
  python -m grpc_tools.protoc -I=/proto-common:/proto $f --python_out=/genproto --grpc_python_out=/genproto
done

for d in $(find /genproto -type d); do
    touch $d/__init__.py
done

echo "import os

os.sys.path.insert(0, os.path.abspath(os.path.dirname(__file__)))" > /genproto/__init__.py
