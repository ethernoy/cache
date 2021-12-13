FROM golang:1.17-alpine

WORKDIR /home/appuser/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN pwd && ls -l  && ls -l /home/appuser
RUN go build -o ./sample-go

FROM alpine:3.14.3  
RUN apk add --no-cache curl jq bash ca-certificates openssl \
    && update-ca-certificates && addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
WORKDIR /home/appuser/app
COPY --from=0 /home/appuser/app/sample-go ./
EXPOSE 8000


CMD [ "./sample-go" ]
