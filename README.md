# QuantChef

A social media for people who like to cook and know the specifics.

QuantChef aims to focus on providing useful information and metadata relating to recipes such as 
accurate recipe times, calorie counts, washing up produced and cost per portion. It also aims to 
make the creation and sharing of recipes easy!

## Development

### Requirements

The app is built using a GO backend and React frontend and for development environments the app is 
run in docker; so in order to develop you must have the correct software installed.

1. Download and install [docker](https://docs.docker.com/engine/install/) and 
[docker compose](https://docs.docker.com/compose/install/)

2. Download and install [GO](https://go.dev/doc/install)

3. Download and install [nodejs](https://nodejs.org/en/download/) (Or install and use 
[NVM](https://github.com/nvm-sh/nvm#installing-and-updating))

4. Install [yarn](https://classic.yarnpkg.com/lang/en/docs/install/#debian-stable)

### MAKE commands

To build the docker image used for development:

`make build`

To start the server on http://localhost:8080

`make up`

To start the frontend in dev mode (for auto-reload on frontend changes etc.) on http://localhost:3000:

`make watch`

#### Linting

To run linting:

`make lint`

* `make lint-fe` 

* `make lint-be`

To run linting with auto formatting:

`make fmt`

* `make fmt-fe` 

* `make fmt-be`
