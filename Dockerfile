# First Stage
ARG GO_VERSION="1.19.1"
FROM golang:${GO_VERSION}-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./item-app

# Second Stage
FROM gcr.io/distroless/base-debian10

WORKDIR /

LABEL maintainer="jinfluenza"

COPY --from=build /app/item-app /item-app

EXPOSE 4040

USER nobody

ENTRYPOINT ["/item-app"]