FROM golang:1.8
LABEL maintainer="thomahl@stud.ntnu.no"
# RUN apt-get update && apt-get -y upgrade
RUN go get -d github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor
RUN go install github.com/tholok97/IMT2681-assignement-2/cmd/currencyMonitor

WORKDIR $GOPATH

CMD ["./bin/currencyMonitor"]
