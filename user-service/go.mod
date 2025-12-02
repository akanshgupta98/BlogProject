module github.com/akanshgupta98/BlogProject/user-service

go 1.24.1

require (
	github.com/akanshgupta98/BlogProject/proto v0.0.0
	github.com/golang-migrate/migrate/v4 v4.19.1
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.74.2
)

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/protobuf v1.36.7 // indirect
)

replace github.com/akanshgupta98/BlogProject/proto => ../proto
