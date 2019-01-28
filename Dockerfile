FROM golang:1.11-alpine as builder

WORKDIR /go/src/github.com/microwaves/go-demoapp
ADD . .
RUN go get && CGO_ENABLED=0 go build -o /demoapp -a -tags netgo -ldflags '-w' main.go

FROM scratch
COPY --from=builder /demoapp /demoapp

ENTRYPOINT ["/demoapp"]
EXPOSE 8080
