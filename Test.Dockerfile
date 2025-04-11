# Use the following commands to run test:
# docker build -f Test.Dockerfile -t kinder-library-test .
# docker run --rm kinder-library-test
FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

CMD ["go", "test", "-v", "./cmd/api"]
