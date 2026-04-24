.PHONY: build

BINARY_NAME=asduptime

build:
	go mod tidy && \
	templ generate && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

dev:
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

clean:
	go clean

docker-dev:
	docker run -v $(PWD):/app -p 8000:8000 -w /app asd-uptime air

docker-build:
	docker build -t asd-uptime .

docker-run:
	docker run -v $(PWD)/data:/app -p 8000:8000 asd-uptime

