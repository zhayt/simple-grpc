mkdir -p ../../pb/user_v1 |
protoc --go_out=../../pb/user_v1 --go_opt=paths=source_relative --go-grpc_out=../../pb/user_v1 --go-grpc_opt=paths=source_relative *.proto

