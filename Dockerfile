FROM golang:1.18-bullseye

RUN go install github.com/CaiqueCastro07/portfolio-backend-go@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 3002
CMD ["go","run","."]