FROM golang:1.24.2-alpine3.20

WORKDIR /app
RUN apk add --no-cache git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN apk add --no-cache postgresql-client
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh
COPY . .
COPY ./migrations /migrations
RUN go build -o main ./
ENTRYPOINT [ "./main" ]