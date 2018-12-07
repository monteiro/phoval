FROM golang:1.11-alpine AS build-stage

RUN apk --no-cache add \
    g++ \
    git \
    make

WORKDIR /src
COPY . .
RUN go build -o bin/phoval-linux cmd/main.go

# Final image.
FROM alpine:latest
RUN apk --no-cache add \
  ca-certificates

# copy the translation templates
COPY --from=build-stage /src/messages/*.txt /usr/local/messages/
COPY --from=build-stage /src/bin/phoval-linux /usr/local/bin/phoval-linux
ENTRYPOINT ["/usr/local/bin/phoval-linux", "-hostdb=mysql", "-template-folder=/usr/local/messages"]
