# Golang
FROM golang:latest as golang-builder
RUN mkdir /go/src/gocoop
WORKDIR /go/src/gocoop
ADD . ./
RUN go get -d -v && go build -o gocoop

# Build
FROM node:latest as angular-builder
LABEL maintainer="francois.allais@hotmail.com"
WORKDIR /app
COPY ./ /app/
RUN npm install
RUN $(npm bin)/ng build --prod --extract-css false

# Final
FROM armhf/ubuntu
LABEL maintainer="francois.allais@hotmail.com"
RUN mkdir -p /app/backend /app/frontend 
COPY --from=golang-builder /go/src/gocoop /app/backend/
COPY --from=angular-builder /app/dist/ /app/frontend/
USER nobody:nobody
WORKDIR /app
ENTRYPOINT [ "/app/backend/gocoop" ]
CMD [ "--help" ]
