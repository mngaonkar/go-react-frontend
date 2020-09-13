FROM golang:alpine

WORKDIR /dist

ADD public ./public
ADD server ./server
COPY main.go .
COPY go.mod .

RUN go mod download
RUN go build main.go

EXPOSE 80

CMD ["/dist/main"]




