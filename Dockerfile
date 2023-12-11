FROM golang:1.19.3

WORKDIR /

COPY ./ /

RUN go get github.com/prisma/prisma-client-go@v0.16.2

RUN make pcg

RUN go mod download

RUN go build -o main main.go

EXPOSE 9000

CMD [ "./main" ]