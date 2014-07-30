NAME=littlegopher
HARDWARE=$(shell uname -m)
VERSION=0.1.0

build:
	mkdir -p stage
	go build -o stage/$(NAME)
	docker build -t $(NAME) .

clean:
	rm -rf stage
	rm -rf $(NAME) 
