FROM golang:1.24.2-alpine3.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main ./
ENTRYPOINT [ "./main" ]