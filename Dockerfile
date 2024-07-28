FROM golang:latest

WORKDIR $GOPATH/src/github.com/vitorbenites/scrapper
COPY . .
RUN go mod download && \
go install

EXPOSE 5000

CMD ["scrapper"]
