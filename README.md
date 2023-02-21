# Implementation Notes
- The proto timestamp field is a string instead of a google.protobuf.Timestamp, this means a conversion has to happen, which means a decision about where to put that conversion has to happen. I recommend using an actual Timestamp in the proto file.
- The only way to effectively test a data store is against the storage infrastructure. I am not familiar enough with liquibase to be able to quickly set up a testcontainer with the db running with the correct, current schema. I also am not willing to submit code with no tests, so I used sqlmock...which is more or less pointless. It has been years since I used sql-mock, I am reminded why this is an utter mistake. I really should have taken the time to figure out how to write a real integration test using liquibase.
- I changed things a little to separate concerns and make testing easier. Maybe for such a small, contrived example as this it is not necessary.
  - Added a store interface/implementation: I want data storage to be separate from message consumption/deserializing. These are entirely separate concerns and shouldn't live in the same functional unit.
  - Added a message processor: This treats the `messages.IntradayValue` type as a sort of domain type. I don't love that, but for the tiny little scope of the application and the time constraints, it'll do. The way it was set up, with the message being processed in an unexported function in the consumer, there was no way to test the behavior of the message processor in isolation.
- I can't find any unique key constraints on the DB...but it seems like there should be? For a given ticker, it seems like we should disallow entries having the same timestamp value. In the case that we did, I would refactor the store to consider THAT unique key constraint violation and return a sentinel error representing that failure mode.
- By the time I got to the HTTP handler, I was in a hurry to finish and probably did not do as complete a job there as I would like. I implemented really basic pagination...and a really basic happy path testcase that at least tests the handler with its pagination middleware. I believe it needs more test cases but just ran out of time. I hope what is here demonstrates to you that I have some sense about where and what to test and know how to do it.


# Go Interview Challenge

Template interview challenge for the GPS Insight api team

This repo contains a little demo service for pulling data about intraday stock values off of a kafka stream and inserting them into a postgres database. The service also has a rest api for making that data consumable by clients.


## What is the point of this challenge?

- Demonstrate an ability to quickly get comfortable in a new codebase and make contributions
- Demonstrate familiarity with the technologies we use every day
- Show us your approach to solving problems


### Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)


## What will you be implementing?

- Consume kafka stream and write to postgres database
- Implement REST endpoint to get stock info from the database

![go-interview-challenge visualization](/overview.png "go-interview-challenge visualization")

If you are looking for ways to spice up your project here are some ideas...
- Try using gRPC or GraphQL instead of REST
- Show us your take on testing
- Allow for the data to be queried with filters/pagination/etc.


## How long should this challenge take?

You are welcome to put in as much time as you like but you should be able to finish it within an hour or so


## Getting Started

- Download and unpack zip of repository

<img src="/download.png" alt="download" width="400">

<p></p>

- Search the repo for `TODO:` tags (there should be two) and follow the instructions


### Running locally

To start local development environment with sample data and live reload of go-interview-challenge run
```
make run
```


## Generating data

To generate message on the kafka topic, use the command
```
make generate-messages
```
