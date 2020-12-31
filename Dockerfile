FROM golang:latest as BUILDER
WORKDIR /app/src

# Copy go mod and sum files
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o api .

FROM alpine:latest as DEPLOY
COPY --from=BUILDER /app/src/api .

EXPOSE 9090
ENTRYPOINT ["./api"]