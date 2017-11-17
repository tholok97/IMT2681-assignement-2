# Solution to IMT2681 Cloud Technologies assignement 2 AND assignement 3

## Introduction

This repo contains our group's (Denbestegruppa) solution to assignement 3, as well as my personal solution to assignement 2. The code is written to be run both with Heroku and Docker, and (should) currently be running in Heroku on the link provided below

* **Assignement 2 instructions**: [instructions2.md](./instructions2.md)
* **Assignement 3 instructions**: [instructions3.md](./instructions3.md)
* **The service**: <https://imt2681-assignement-2.herokuapp.com/>

*(This repo is WIP. Explanation and usage-examples will be added as the repo matures)*

## Installation

1. Get the repo onto your system with `go get -d github.com/tholok97/IMT2681-assignement-2/`. (-d tells go get to not try and build anything).
2. Install the repo with `go install ./cmd/subscriberWebAPI/ && go install ./cmd/currencyMonitor/`.
3. To run the repo with Heroku / Docker you need to supply your own '.env' file. Make a file called '**.env**' that looks like this: 

        PORT=5000
        SCHEDULE_HOUR=00
        SCHEDULE_MINUTE=10
        SCHEDULE_SECOND=00
        FIXER_IO_URL=http://api.fixer.io
        MONGO_DB_URL=mongodb://localhost
        MONGO_DB_DATABASE_NAME=assignement_2

## Running with Heroku
        
The command `heroku local` should start running the project locally. If everything went okay you'll see output like this: 

        [OKAY] Loaded ENV .env File as KEY=VALUE Format
        20:49:50 web.1     |  Read env  PORT  =  5000
        20:49:50 web.1     |  Read env  FIXER_IO_URL  =  http://api.fixer.io
        20:49:50 web.1     |  Read env  MONGO_DB_URL  =  mongodb://localhost
        20:49:50 web.1     |  Read env  MONGO_DB_DATABASE_NAME  =  assignement_2
        20:49:50 clock.1   |  Read env  SCHEDULE_HOUR  =  00
        20:49:50 clock.1   |  Read env  SCHEDULE_MINUTE  =  10
        20:49:50 clock.1   |  Read env  SCHEDULE_SECOND  =  00
        20:49:50 clock.1   |  Read env  FIXER_IO_URL  =  http://api.fixer.io
        20:49:50 clock.1   |  Read env  MONGO_DB_URL  =  mongodb://localhost
        20:49:50 clock.1   |  Read env  MONGO_DB_DATABASE_NAME  =  assignement_2
        20:49:50 clock.1   |  Updating monitor...
        20:49:50 web.1     |  Listening on port 5000...
        20:49:51 clock.1   |  Notifying all subscribers...
        20:49:51 clock.1   |  Sleeping  3h20m8.692749179s ...

## Running with Docker

The command `docker-compose up` should build / fetch the necessary containers, link them and start running the project. See the '**docker-compose.yml**' for details.

## Dependencies

* **MongoDB**. To run the code locally you need to have mongod running. To run the code on heroku you need to supply a URI to a mongoDB service as a config variable (ex. mlabs).
* **Heroku** / **Docker**. The project can be run with either Heroku or Docker. It's possible to use neither of these, but you'll still have to inject the environment variables the project expects somehow (see the variables defined in the '**.env**' file above)
* *(mgo.v2)*. The code relies on the mongoDB driver mgo.v2, but It's included in the repo, so you shouldn't have to worry about that.
