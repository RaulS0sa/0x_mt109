
FROM golang:1.15.3
ADD . /go/src/ClarkWrite
WORKDIR /go/src/ClarkWrite
RUN go get ClarkWrite
RUN go install
ENTRYPOINT ["/go/bin/ClarkWrite"]
