FROM golang:1.16.2 
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE  9095
CMD ["/app/main"]


