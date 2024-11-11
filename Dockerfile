FROM ubuntu:latest
LABEL maintainer="jeneksero@gmail.com"
RUN apt-get update -y
FROM golang:1.23.1-alpine as builder
WORKDIR /MyMoney
COPY ["./go.mod", "./go.sum", "./"]
RUN  go mod download
COPY . ./
ENTRYPOINT ["go", "run"]
CMD ["main.go"]