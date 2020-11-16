protoc --go_out=plugins=grpc:./golang/TalkRpc chat.proto
python -m grpc.tools.protoc -I. --python_out=../client/python/lib/api --grpc_python_out=../client/python/lib/api chat.proto

protoc -I=. chat.proto --js_out=import_style=commonjs:../client/web --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:../client/web --plugin=protoc-gen-grpc-web="protoc-gen-grpc-web-1.2.1-windows-x86_64.exe"

git add golang/TalkRPC/chat.pb.go
git add ../client/python/lib/api/
git commit -m "auto: grpc generated commit"
