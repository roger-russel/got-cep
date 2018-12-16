FROM golang:1.11

RUN go get github.com/revel/cmd/revel

RUN revel new got-cep

CMD ["sleep", "infinity"]
