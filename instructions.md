# Assignment 2 - Persistence, Live feed, Notifications

## Instructions

Develop a service, that will allow a user to monitor a currency ticker, and notify a webhook upon certain conditions are met, such as the price falling below or going above a given threshold. The API must allow the user to specify the base currency, the target currency, and the min and max price for the event to trigger the notification. The notification will be provided via a webhook specified by the user, and multiple webhooks should be provided (predefined types). 

In addition, the service will be able to monitor the currencies (all from the http://api.fixer.io/latest?base=EUR query) at regular time intervals (once a day) and store the results in a MongoDB database. The system will allow the user to query for the "latest" ticker of given currency pair between EUR/xxx, and also, to query for the "running average" of the last 7 days. 

## Service specification

### Registration of new webhook

New webhooks can be registered using POST requests with the following schema. Note we will use /root as a placeholder for the root path of your web service (i.e. the path you will submit to the submission system). For example, if your web service runs on https://localhost:8080/exchange, then this is the root path you would submit.

#### Request

Path: /root

Payload specification:

```
{

    "webhookURL": {
        "type": "string"
    },

        "baseCurrency": {
            "type": "string"
        },
        "targetCurrency": {

            "type": "string"
        },

        "minTriggerValue": {
            "type": "number"
        }, 

        "maxTriggerValue": {
            "type": "number"
        }
}
```


Example:

```
{

    "webhookURL": "http://remoteUrl:8080/randomWebhookPath",

        "baseCurrency": "EUR",
        "targetCurrency": "NOK",

        "minTriggerValue": 1.50, 

        "maxTriggerValue": 2.55
}
```

#### Response

The response body should contain the id of the created resource (only the created id, not the entire path; no json encoding).

### Invoking a registered webhook

When invoking a registered webhook, use the following payload specification:

#### Request

Path: webhookUrl

Payload Specification:

```
{

    "baseCurrency": {
        "type": "string"
    },
        "targetCurrency": {

            "type": "string"
        },
        "currentRate": {
            "type": "number"
        },

        "minTriggerValue": {
            "type": "number"
        }, 

        "maxTriggerValue": {
            "type": "number"
        }
}
```

Example:

```
{

    "baseCurrency": "EUR",
        "targetCurrency": "NOK",
        "currentRate": 2.75,

        "minTriggerValue": 1.50, 

        "maxTriggerValue": 2.55
}
```

#### Response

Upon successful notification you will receive either status code 200 or 204.

### Accessing registered webhooks

Registered webhooks should be accessible using the GET method and the webhook id generated during registration.

#### Request

Path: /root/{id}

#### Response

Body: (see POST request)

### Deleting registered webhooks

Registered webhooks can further be deleted using the DELETE method and the webhook id.

Path: /root/{id}

### Retrieving the latest currency exchange rates

#### Request

Path:  /root/latest

Body:

```
{
    "baseCurrency": "EUR",
        "targetCurrency": "NOK"

}
```

#### Response

The response should contain only the latest exchange rate value (no json tags).

Body: value

Example: 1.56

### Retrieving the running average over the past seven days

#### Request

Path: /root/average

Body:

```
{
    "baseCurrency": "EUR",
        "targetCurrency": "NOK"

}
```

#### Response

The response should contain only the running average value (no json tags).

Body: value

Example: 2.75

*Important: For all requests, ensure that you use appropriate status codes and semantics (see IETF RFC 7231 for details).*

## Formal aspects

This assignment is worth 15% of your total mark for the internal portfolio. The assignment is individual. Code snippets used from the web and other alternative sources (StackOverflow, tutorials) must be clearly attributed to the original source, in the source code. Your code needs to have a test coverage of at least 40%. Note that your tests actually need to pass to meet this requirement. There is +3 points bounty for extra work.

*Due date: 2nd November 2017, Thursday, 7am*

Submission system will be available at http://prod3.imt.hig.no/submission/



## Resources

* http://fixer.io/  Currency ticker data
* https://mlab.com/  MongoDB cloud hosting (choose the FREE plan)
* http://gopkg.in/mgo.v2 MongoDB Go driver
* https://discordapp.com/developers/docs/resources/webhook Discord webhooks


Note: most likely we will use Heroku-based deployment, but, we need to test it first.

