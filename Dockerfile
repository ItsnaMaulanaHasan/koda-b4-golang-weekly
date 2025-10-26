FROM golang:tip-alpine3.22

WORKDIR /app

COPY go.* ./

RUN go mod download || true

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]