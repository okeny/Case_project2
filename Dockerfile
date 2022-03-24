FROM golang:latest

LABEL maintainer="okeny <okeny.patrick@andela.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 3000

RUN go build

CMD ["./project-cases"]