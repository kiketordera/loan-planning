FROM golang:alpine as builder

ENV GIN_MODE=release

RUN apk update && apk add git && apk add ca-certificates

COPY . .
WORKDIR /

RUN go get -d -v $GOPATH/src/github.com/kiketordera/loan-planning


 RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/loan-planning $GOPATH/src/github.com/kiketordera/loan-planning

FROM scratch
COPY --from=builder /go/bin/loan-planning /loan-planning

EXPOSE 8080/tcp

ENV GOPATH /go
ENTRYPOINT ["/loan-planning"]

