# Build
FROM node:latest as builder
LABEL maintainer="francois.allais@hotmail.com"
WORKDIR /app
COPY ./ /app/
RUN npm install
RUN $(npm bin)/ng build --prod --extract-css false

# Caddy
FROM elswork/arm-caddy
LABEL maintainer="francois.allais@hotmail.com"
COPY --from=build /app/dist/watcher-frontend/ /srv/
COPY ./Caddyfile /etc/Caddyfile
EXPOSE 2015
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout"]