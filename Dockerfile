FROM golang:1.14-alpine AS base
RUN apk add --no-cache git
WORKDIR /app


FROM base as builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o unleash-checkr main.go


FROM scratch as final
WORKDIR /app
COPY --from=builder /app/unleash-checkr /app
ENTRYPOINT ["./unleash-checkr"]