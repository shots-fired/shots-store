# Shots Store

[![Build Status](https://travis-ci.org/shots-fired/shots-store.svg?branch=master&service=github)](https://travis-ci.org/shots-fired/shots-store)
[![Coverage Status](https://coveralls.io/repos/github/shots-fired/shots-store/badge.svg?branch=master&service=github)](https://coveralls.io/github/shots-fired/shots-store?branch=master)

Shots is a Discord bot. This project houses all the code responsible for updating the bot's database(s).

## Contributing

1. Install Go 1.11 or higher
2. `go get -u github.com/shots-fired/shots-store`

## Running

The easiest way to run Shots is by cloning the `shots-deploy` project and using Docker Compose.

1. `git clone github.com/shots-fired/shots-deploy`
2. `cd shots-deploy`
3. `docker-compose build`
4. `docker-compose up`

## Environment variables

* SERVER_ADDRESS
* REDIS_ADDRESS
* REDIS_PASSWORD
