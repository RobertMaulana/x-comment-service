FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/comment-service
COPY . .
RUN go get
RUN go build -o comment-service
ENTRYPOINT ./comment-service

# running in port
EXPOSE 8080

# expose GRPC port
EXPOSE 6060
