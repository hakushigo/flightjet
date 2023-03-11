FROM golang:latest
WORKDIR /url
COPY . .
RUN go build -o main main.go
EXPOSE 3000
CMD ["./main"]