# Go Interview Challenge

Template interview challenge for the GPS Insight api team

This repo contains a little demo service for pulling data about intraday stock values off of a kafka stream and inserting them into a postgres database. The service also has a rest api for making that data consumable by clients.


## What is the point of this challenge?

- Demonstrate an ability to quickly get comfortable in a new codebase and make contributions
- Demonstrate familiarity with or ability to pick up the technologies we use every day
- Show us your approach to solving problems


### Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)


## What will you be implementing?

- Consume kafka stream and write to postgres database
- Implement REST endpoint to get stock info from the database

![go-interview-challenge visualization](/overview.png "go-interview-challenge visualization")

These basic requirements are designed to be relatively simple and straightforward to implement. If that is all you have time for that is totally fine. It will give us a flavor for what you can do. However, if you want to go beyond the basics and are looking for ways to spice up your project here are some ideas...
- Allow for the data to be queried with filters/pagination/etc.
- Refactor the project to show how you like to organize your code
- Show us your take on testing
- Get creative! This is your chance to make your project stand out. Have fun!


## Getting Started

- Download and unpack zip of repository

<img src="/download.png" alt="download" width="400">

<p></p>

- Search the repo for `TODO:` tags (there should be four) and follow the instructions


### Running locally

To start local development environment with sample data and live reload of go-interview-challenge run
```
make run
```


## Generating data

To generate message on the kafka topic, open a separate terminal session and use the command
```
make generate-messages
```
