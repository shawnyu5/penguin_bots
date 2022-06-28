FROM golang:1.18.3-alpine3.16 AS build

WORKDIR /server

# copy server files
COPY ./go.* ./
COPY ./server.go ./

COPY ./utils/ ./utils/
COPY ./coin_products/ ./coin_products/

# RUN ls

RUN go get github.com/joho/godotenv
RUN go get github.com/patrickmn/go-cache
RUN go build -o server

FROM golang:1.18.3-alpine3.16 AS production

LABEL maintainer="Shawn Yu"
LABEL description="A golang web server for penguin bot"

ENV PORT=8080

WORKDIR /server
COPY --from=build /server/server ./server
COPY ./.env ./

CMD ./server

EXPOSE 8080
