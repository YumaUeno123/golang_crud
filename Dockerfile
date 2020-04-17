FROM golang:1.14
RUN mkdir /go/src/first_docker_project
WORKDIR /go/src/first_docker_project
ADD . /go/src/first_docker_project