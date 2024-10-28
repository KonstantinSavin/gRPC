prot:
		protoc -I protos/proto/sso --go_out=./protos/gen/go/ --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go --go-grpc_opt=paths=source_relative protos/proto/sso/sso.proto

migration:
		go run ./sso/cmd/migrator --storage-path=./sso/storage/sso.db --migrations-path=./sso/migrations
