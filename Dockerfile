# First Stage
ARG GO_VERSION="1.19.1"
FROM alpine:${GO_VERSION}-buster AS build

LABEL maintainer="jinfluenza"

WORKDIR /app

COPY * ./

RUN go mod download

RUN go build -o ./item-app

# Second Stage
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/item-app /item-app

EXPOSE 4040

USER nobody

ENTRYPOINT ["/item-app"]