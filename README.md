# Mock API Server

## Quick Start

1. Setup Golang and NodeJS.
1. Run `npm install -g yarn` to get the package manager yarn.
1. Setup an instance of MongoDB. You can either use one running on your localhost or one in the cloud.
1. Setup the environment variables using `example.env`. The logic for this sits in `internal/env.go`.
1. Run `make dev-be` to expose the server on port 3080.
1. Run `make dev-fe` to run frontend.

## File Structure

`/cmd`

Stores the executables like the server.

`/domain`

Stores all the interfaces and DB models used in the application. The DB schema of the application can be inferred from the code here.

`/internal`

Utilises and implements the interfaces in `/domain`. See the README in `/internal` for more details.

`/web`

Holds the code for frontend, which is written in React. See the README in `/web` for more details.

## Tech Stack

FE: React
BE: Golang, with `go-chi` and `logrus` being notable dependencies
DB: MongoDB

## Getting an instance of MongoDB

Local Instance:

- <https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/>
