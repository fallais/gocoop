# Build (Golang)
FROM golang:latest as builder-golang
WORKDIR /go/src/gocoop
ADD . /go/src/gocoop
# RUN env GOOS=linux GOARCH=arm GOARM=7
RUN go get -d -v ./...
RUN go build -o /go/bin/gocoop

# Build (Angular)
FROM node:latest as builder-angular
LABEL maintainer="francois.allais@hotmail.com"
WORKDIR /app
COPY ./web/ /app/
RUN npm install
RUN $(npm bin)/ng build --prod --extract-css false

# Ubuntu
FROM ubuntu:latest
LABEL maintainer="francois.allais@hotmail.com"
COPY --from=builder-golang /go/bin/gocoop /usr/bin/gocoop
COPY --from=builder-angular /app/dist/gocoop/ /srv/
COPY entrypoint.sh /entrypoint.sh
COPY Caddyfile /etc/Caddyfile
RUN apt-get update && \
    apt-get install -y curl && \
    curl https://getcaddy.com | bash -s personal && \
    chmod +x /entrypoint.sh
WORKDIR /
ENTRYPOINT ["/entrypoint.sh"]