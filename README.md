# My solution to assignement 2 in IMT2681 Cloud Technologies

Links:

* **Assignement instructions**: [instructions.md](./instructions.md)
* **The service**: <https://imt2681-assignement-2.herokuapp.com/>

*(This repo is WIP. Explanation and usage-examples will be added as the repo matures)*

## Installation

1. Get the repo onto your system with `go get -d github.com/tholok97/IMT2681-assignement-2/`. (-d tells go get to not try and build anything).
2. Install the repo with `go install ./cmd/subscriberWebAPI/ && go install ./cmd/currencyMonitor/`.
3. To run the repo with Heroku you need to supply your own '.env' file. Make a file called '.env' that looks like this: 
    PORT=5000
    SCHEDULE_HOUR=00
    SCHEDULE_MINUTE=10
    SCHEDULE_SECOND=00
    FIXER_IO_URL="http://api.fixer.io"
    MONGO_DB_URL="mongodb://localhost"
    MONGO_DB_DATABASE_NAME="assignement_2"
4. The command `heroku local` should now start running the project locally. If everything went okay you'll see output like this: 
    [OKAY] Loaded ENV .env File as KEY=VALUE Format
    20:40:08 web.1     |  Read env  PORT  =  5000
    20:40:09 web.1     |  Read env  FIXER_IO_URL  =  http://api.fixer.io
    20:40:09 web.1     |  Read env  MONGO_DB_URL  =  mongodb://localhost
    20:40:09 web.1     |  Read env  MONGO_DB_DATABASE_NAME  =  assignement_2
    20:40:09 clock.1   |  Read env  SCHEDULE_HOUR  =  00
    20:40:09 clock.1   |  Read env  SCHEDULE_MINUTE  =  10
    20:40:09 clock.1   |  Read env  SCHEDULE_SECOND  =  00
    20:40:09 clock.1   |  Read env  FIXER_IO_URL  =  http://api.fixer.io
    20:40:09 clock.1   |  Read env  MONGO_DB_URL  =  mongodb://localhost
    20:40:09 clock.1   |  Read env  MONGO_DB_DATABASE_NAME  =  assignement_2
    20:40:09 clock.1   |  Updating monitor...
    20:40:09 web.1     |  Listening on port 5000...
    20:40:10 clock.1   |  Notifying all subscribers...
    20:40:10 clock.1   |  Notifying  https://requestb.in/z4jlg3z4
    20:40:10 clock.1   |    Rate:  9.5923
    20:40:10 clock.1   |    MinTriggerValue:  1.5
    20:40:10 clock.1   |    MaxTriggerValue:  2.55
    20:40:10 clock.1   |    done trying
    20:40:10 clock.1   |  Sleeping  3h29m49.650031455s ...

## Dependencies

* **MongoDB**. To run the code locally you need to have mongod running. To run the code on heroku you need to supply a URI to a mongoDB service (for example mlabs).
* **Heroku**. The code is written with Heroku in mind, so the easiest way to run it is using Heroku commands (see step 3). It's possible to do without heroku, but you'll have to supply the environment-variables defined in the .env file in some other way.
* (mgo.v2). The code relies on the mongoDB driver mgo.v2, but It's included in the repo, so you shouldn't have to worry about that.
