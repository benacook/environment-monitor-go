FROM golang:1.13-alpine

WORKDIR /src/environment-monitor-go

COPY . .

RUN go get -d -v ./...
RUN go install ./...

EXPOSE 3000

ENTRYPOINT ["environment-monitor-go"]
