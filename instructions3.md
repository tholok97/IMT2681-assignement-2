# Instructions

This assignment, Assignment 3, has two aspects, and builds on the services developed in Assignment 2.

## Aspect 1: 

Develop a service that allows users to interact with a bot on a instant messaging system of your choice (e.g. Slack), and provide currency conversions. The bot will form an interactive user-interface that relies on the currency conversion service that you have built in Assignment 2.

Note that the messaging system needs to be able to support both incoming and outgoing webhooks without any additional layer of complexity (e.g. authentication). Check https://dialogflow.com/docs/integrations/ for details.

## Aspect 2: 

Assignment 3 functionality must be deployed on mLab (for database functionality) and on Heroku (for computational functionality) and on Dialogflow (for the bot integration). However, you are also required to prepare Dockerfiles and docker-compose configuration (for your code & MongoDB), such that the solution to Assignment 3 could potentially be re-deployed on an alternative cloud provider that supports Docker containers, i.e. not relying on mLabs or on Heroku. 

# Service specification

## Bot

The bot should be able to answer simple questions about the current currency conversion rates. Eg.

"What is the current exchange rate between Norwegian Kroner and Euro?"

"What is the exchange rate between USD and NOK?"

"What is the exchange rate between euro and norwegian kroner?"

### Response

The bot should respond in English, with the current exchange rates. 
