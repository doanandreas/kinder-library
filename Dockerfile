FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./cmd/api/... -v

RUN go build -o /kinder-library ./cmd/api

EXPOSE 8080

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

ENTRYPOINT ["/wait-for-it.sh", "postgres:5432", "--", "/kinder-library"]

# You can use this target to run tests
# docker build --target test -t kinder-library-test .
# docker run --rm kinder-library-test
FROM golang:1.24.2 AS test

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

CMD ["go", "test", "-v", "./cmd/api"]