FROM golang:latest as BUILDER
WORKDIR /app/src

# Copy go mod and sum files
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -a -v -o app .

FROM golang:alpine as DEPLOY
WORKDIR /app/build
COPY --from=BUILDER app .

CMD ["./app"]