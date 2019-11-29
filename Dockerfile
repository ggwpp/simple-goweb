### build stage ###
FROM golang:1.11.5-alpine as build
LABEL maintainer="patipat@omise.co"

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

### Add App stage ###
FROM alpine

WORKDIR /app
COPY --from=build /go/src/app/index.html .
COPY --from=build /go/src/app/app .
ENTRYPOINT ["/app/app"]
EXPOSE 3000

