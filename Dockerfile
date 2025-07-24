FROM golang:1.24-alpine AS build

COPY ./ /go/src/ 
RUN go build -C /go/src/.

FROM alpine:latest AS release

COPY --from=build /go/src/bumpy /usr/bin/bumpy

ENTRYPOINT ["bumpy","server"]
