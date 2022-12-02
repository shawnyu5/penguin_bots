FROM golang:1.18.3-alpine3.16 AS build

WORKDIR /logger

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./logger.go ./

RUN go build -o logger

FROM golang:1.18.3-alpine3.16 AS production

WORKDIR /logger
RUN apk --no-cache add gcc g++
COPY --from=build /logger/logger ./
# COPY ./.env ./

CMD ["./logger"]
