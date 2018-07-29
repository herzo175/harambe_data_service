NAME=harambe_user_service
VERSION=$(shell git rev-parse HEAD)
DEFAULT_PORT=3000

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o harambe_data_service

build:
	docker build --build-arg PORT_ARG=$(DEFAULT_PORT) -t $(NAME)/$(VERSION) ./

run:
	docker run --rm -it \
		-p $(DEFAULT_PORT):$(DEFAULT_PORT) \
		$(NAME)/$(VERSION)

clean:
	docker stop $(NAME):$(VERSION)