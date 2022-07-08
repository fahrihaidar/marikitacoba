FROM golang:1.18

WORKDIR /home

COPY . .

RUN go get -u 

RUN go build -o main .

CMD ["./main"]