FROM golang:1.13 as builder

WORKDIR /go/src/github.com/jeonjonghyeok/golang
RUN go get -d -v github.com/urfave/cli

# COPY main.go .
# RUN GOOS=linux go build -a -o greet .

FROM busybox
WORKDIR /opt/greet/bin

# COPY --from=builder /go/src/github.com/jeonjonghyeok/golang/ .
# ENTRYPOINT ["./greet"]
