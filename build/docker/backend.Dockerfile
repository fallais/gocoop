# Build
FROM golang:latest as builder
RUN mkdir /go/src/gocoop
WORKDIR /go/src/gocoop
ADD . ./
RUN go get -d -v && go build -o gocoop

# Ubuntu
FROM armhf/ubuntu
LABEL maintainer="francois.allais@hotmail.com"
RUN mkdir -p /app/backend /app/frontend 
COPY --from=builder /go/src/gocoop /usr/bin/gocoop
USER nobody:nobody
WORKDIR /
ENTRYPOINT [ "/usr/bin/gocoop" ]
CMD [ "--help" ]
