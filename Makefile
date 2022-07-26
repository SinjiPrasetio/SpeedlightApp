BINARYNAME=Speedlight.exe

build:
	@echo "Updating Speedlight..."
	@go get github.com/SinjiPrasetio/speedlight
	@go mod vendor
	@go mod tidy
	@echo "Building Speedlight..."
	@go build -o tmp/${BINARYNAME}
	@echo "Speedlight has been built successfully."

run: build
	@echo "Starting Speedlight..."
	@./tmp/${BINARYNAME}
	@echo "Speedlight started."

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARYNAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./..
	@echo "Testing Done!"

start: run

stop:
	@echo "Stopping Speedlight..."
	@-pkill -SIGTERM -f "./tmp/${BINARYNAME}"
	@echo "Speedlight has been stopped successfully."

restart: 
	@echo "Restarting Speedlight..."
	@make stop
	@make start
	@echo "Speedlight has been restarted successfully."