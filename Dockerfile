FROM golang:1.15

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
RUN go build main.go

CMD ["/app/main"]