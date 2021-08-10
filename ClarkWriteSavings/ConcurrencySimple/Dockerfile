## We specify the base image we need for our
## go application
FROM golang:1.15.3
## We create an /app directory within our
## image that will hold our application source
## files
#RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
#ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
#WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
#RUN go build -o pixelremover .
## Our start command which kicks off
## our newly created binary executable
#ENTRYPOINT ["go", "run" ,"pixelremover", "test_1"]
#CMD ["pixelremover"]

ADD . /go/src/pixelremover
WORKDIR /go/src/pixelremover
RUN go get pixelremover
RUN go install
ENTRYPOINT ["/go/bin/pixelremover"]
