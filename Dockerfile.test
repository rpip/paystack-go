FROM golang:1.7

RUN mkdir -p /go/src/github.com/rpip/paystack-go
WORKDIR /go/src/github.com/rpip/paystack-go
ENV GOPATH /go

RUN go install -race std && go get golang.org/x/tools/cmd/cover
RUN go get github.com/golang/lint/golint github.com/Masterminds/glide

COPY . /go/src/github.com/rpip/paystack-go

CMD ["./runtests.sh"]
