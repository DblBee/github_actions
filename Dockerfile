FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD . /app

ENV PORT 8080

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Command to run the executable
CMD ["./main"] 