# Dockerfile for currencyMonitor container.


FROM golang:1.8
LABEL maintainer="thomahl@stud.ntnu.no"

# update and upgrade for security reasons
RUN apt-get update && apt-get -y upgrade

# make all subsequent copies relative to $GOTPATH
WORKDIR $GOPATH

# copy over relevant files from the LOCAL repository
COPY ./cmd/currencyMonitor ./src/github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor
COPY ./currencyWebhookService ./src/github.com/tholok97/IMT2681-assignement-2/currencyWebhookService
COPY ./vendor ./src/github.com/tholok97/IMT2681-assignement-2/vendor

# make binary
RUN go install github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor

# run binary
CMD ["./bin/currencyMonitor"]
