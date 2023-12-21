FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8888

RUN go build  -o /out/main ./

ENTRYPOINT ["/out/main"]

