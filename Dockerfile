FROM golang:1.20 as builder
ARG VERSION

WORKDIR /app

COPY . /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o bassoon cmd/app/main.go

FROM gcr.io/distroless/static-debian10
COPY --from=builder /app/bassoon /bin/bassoon

ENTRYPOINT ["/bin/bassoon"]