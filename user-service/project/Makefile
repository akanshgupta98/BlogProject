CMD_PATH=../cmd
run: clean
	mkdir -p bin
	go build -o ../bin/user-service ${CMD_PATH}/

clean: 
	rm -rf ../bin

migrate-create:
	migrate create -ext sql -dir ${CMD_PATH}/migrate/migrations -seq $(seq)

migrate-clean:
	rm -rf ${CMD_PATH}/migrate/migrations

migrate-up:
	go run ${CMD_PATH}/migrate/main.go --oper=up

migrate-down:
	go run ${CMD_PATH}/migrate/main.go --oper=down
