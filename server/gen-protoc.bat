protoc --go_out=plugins=grpc:./golang/TalkRpc chat.proto
python -m grpc.tools.protoc -I. --python_out=../client/python/lib/api --grpc_python_out=../client/python/lib/api chat.proto

git add golang/TalkRPC/chat.pb.go
git add ../client/python/lib/api/
git commit -m "auto: grpc generated commit"
