FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .

RUN go build -o /kinder-library

EXPOSE 8080

CMD [ "/kinder-library" ]