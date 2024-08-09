FROM golang:latest

WORKDIR .

COPY . .

RUN go mod download

RUN go build -o server .

EXPOSE 4000

CMD ["./server"]
