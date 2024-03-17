# needed for bash in sqlc image
FROM busybox:1.35.0-uclibc as busybox

# generate sqlc
FROM sqlc/sqlc:latest AS sql_builder
ADD . /app
WORKDIR /app
COPY --from=busybox /bin/sh /bin/sh
ENTRYPOINT ["/bin/sh"]
RUN ["/workspace/sqlc", "generate"]

# backend
FROM golang:alpine AS builder
ADD . /app
WORKDIR /app
COPY --from=sql_builder /app/generated generated
RUN apk add --no-cache build-base
RUN go mod download
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -a -o dist/app cmd/app/app.go
RUN ldd dist/app

# Container to deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates musl-utils bash
RUN apk --no-cache add sqlite-libs sqlite-dev
COPY --from=builder /app/dist/app /app
ENV PORT 8080
EXPOSE 8080
RUN chmod +x /app
ENTRYPOINT [ "/bin/bash", "-l", "-c" ]
CMD ["./app"]
