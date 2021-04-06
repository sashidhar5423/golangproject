FROM golang:1.16.2
WORKDIR /app
COPY go.mod go.sum
RUN go mod dowload
COPY . .
RUN go build -o main .
EXPOSE 7071
CMD [ "./main" ]


