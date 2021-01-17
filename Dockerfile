FROM golang:1.15-alpine

WORKDIR /todo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]