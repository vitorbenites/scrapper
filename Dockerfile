FROM golang:latest

WORKDIR /app
COPY * ./
RUN go mod download && \
CGO_ENABLED=0 GOOS=linux go build -o /scrapper

EXPOSE 5000

CMD ["/scrapper"]
