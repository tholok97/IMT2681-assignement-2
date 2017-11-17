FROM golang:1.8
LABEL maintainer="thomahl@stud.ntnu.no"

RUN apt-get update && apt-get -y upgrade

WORKDIR $GOPATH

COPY ./cmd/currencyMonitor ./src/github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor
COPY ./currencyWebhookService ./src/github.com/tholok97/IMT2681-assignement-2/currencyWebhookService
COPY ./vendor ./src/github.com/tholok97/IMT2681-assignement-2/vendor

RUN go install github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor


CMD ["./bin/currencyMonitor"]
