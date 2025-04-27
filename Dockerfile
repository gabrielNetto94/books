FROM golang:latest AS builder
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./out


FROM scratch
WORKDIR /app
COPY --from=builder /build/out ./out
EXPOSE 3000
ENTRYPOINT ["./out"]