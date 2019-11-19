FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 1080 1081
CMD ["./main"]