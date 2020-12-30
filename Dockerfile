FROM golang:latest as BUILDER
WORKDIR /app/src

# Copy go mod and sum files
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -v -o api .

FROM alpine:latest as DEPLOY
WORKDIR /app/build
COPY --from=BUILDER app .

EXPOSE 9090
CMD ["./api"]