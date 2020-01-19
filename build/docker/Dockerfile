# Build (Golang)
FROM golang as builder-golang
WORKDIR /go/src/gocoop
ADD . /go/src/gocoop
RUN env GOOS=linux GOARCH=arm GOARM=7
RUN go get -d -v ./...
RUN go build -o /go/bin/gocoop

# Build (Angular)
FROM node:latest as builder-angular
LABEL maintainer="francois.allais@hotmail.com"
WORKDIR /app
COPY ./ /app/
RUN npm install
RUN $(npm bin)/ng build --prod --extract-css false

# Ubuntu
FROM arm32v7/ubuntu
LABEL maintainer="francois.allais@hotmail.com"
RUN mkdir -p /app
COPY --from=builder-golang /go/src/gocoop /usr/bin/gocoop
COPY --from=builder-angular /app/dist/gocoop/ /srv/
RUN apt-get update && \
    apt-get install -y supervisor && \
    curl https://getcaddy.com | bash -s personal
COPY entrypoint.sh entrypoint.sh
COPY Caddyfile /etc/Caddyfile
CMD ./entrypoint.sh