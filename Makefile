NAME=harambe_user_service
VERSION=$(shell git rev-parse HEAD)
DEFAULT_PORT=3000

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/harambe_data_service

build:
	docker build --build-arg PORT_ARG=$(DEFAULT_PORT) -t $(NAME)/$(VERSION) ./

run:
	docker run --rm -it \
		-p $(DEFAULT_PORT):$(DEFAULT_PORT) \
		$(NAME)/$(VERSION)

deploy:
	# TODO: move to CI platform (probably Travis CI)
	heroku container:push web
	heroku container:release web

clean:
	docker stop $(NAME):$(VERSION)