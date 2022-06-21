FROM golang:1.18.3-alpine3.16

LABEL maintainer="Shawn Yu"
LABEL description="A golang web server for penguin bot"

ENV PORT=8080

RUN mkdir sever
RUN mkdir utils
RUN mkdir coin_products

WORKDIR /coin_products
COPY ../coin_products/go.mod ./
RUN ls

# WORKDIR /server
# # copy over server files
# COPY server.go server.go
# COPY go.mod go.mod
# COPY go.sum go.sum

# RUN go build .

# RUN ls

EXPOSE 8080
