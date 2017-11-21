# Dockerfile for subscriberWebAPI container.

#------------------------------ BUILDER CONTAINER ------------------------------

FROM golang:1.8 as builder
LABEL maintainer="thomahl@stud.ntnu.no"

# update and upgrade for security reasons
RUN apt-get update && apt-get -y upgrade

# make all subsequent copies relative to $GOTPATH
WORKDIR $GOPATH

# copy over relevant files from the LOCAL repository
COPY ./cmd/subscriberWebAPI ./src/github.com/tholok97/IMT2681-assignement-2-and-3/cmd/subscriberWebAPI
COPY ./currencyWebhookService ./src/github.com/tholok97/IMT2681-assignement-2-and-3/currencyWebhookService
COPY ./vendor ./src/github.com/tholok97/IMT2681-assignement-2-and-3/vendor

# move into the directory where the 'main.go' file is located
WORKDIR $GOPATH/src/github.com/tholok97/IMT2681-assignement-2-and-3/cmd/subscriberWebAPI

# create independent binary 
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o subscriberWebAPI

#------------------------------ ACTUAL CONTAINER -------------------------------

FROM scratch
LABEL maintainer="thomahl@stud.ntnu.no"

WORKDIR /

# copy over binary from builder
COPY --from=builder /go/src/github.com/tholok97/IMT2681-assignement-2-and-3/cmd/subscriberWebAPI/subscriberWebAPI .

# copy over ca-certificates
ADD ./ca-certificates.crt /etc/ssl/certs/

# run binary
CMD ["./subscriberWebAPI"]
