# Makefile



# protoc -Iproto --go_out=pb --go-grpc_out=pb proto/profile.proto


# cd proto && protoc -Iproto --go_out=../protogen/golang --go_opt=paths=source_relative \
# --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
# --grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
# --grpc-gateway_opt generate_unbound_methods=true \
# ./user/*.proto

user:
	cd proto && protoc -I. -I./user --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./user/*.proto

profile:
	cd proto && protoc -I. -I./profile --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./profile/*.proto

	

auth:
	cd proto && protoc -Iproto --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./auth/*.proto