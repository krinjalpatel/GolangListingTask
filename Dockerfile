FROM golang:onbuild
FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
ADD . /go/src/github.com/golang/example/outyet
RUN go install github.com/golang/example/outyet
RUN go install github.com/gorilla/mux
RUN go build -o main . 
CMD ["/api/main"]
EXPOSE 8010