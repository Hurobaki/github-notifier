FROM golang:alpine
ADD . /go/src/github.com/Hurobaki/github-notifier
RUN go install github.com/Hurobaki/github-notifier
CMD ["/go/bin/github-notifier"]
EXPOSE 3000
