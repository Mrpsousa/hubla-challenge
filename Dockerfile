FROM golang:latest

WORKDIR /app

COPY . .

EXPOSE 8000

RUN GOOS=linux go build -o server ./

CMD ["./server"]