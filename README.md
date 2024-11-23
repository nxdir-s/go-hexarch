# Go HexArch

This repository serves as a template and documentation on how to implement Hexagonal Architecture within a golang application

### Table of Contents

- [Hexagonal Architecture](#hexagonal-architecture)
- [Project Structure](#project-structure)
  - [cmd](#cmd)
  - [adapters](#adapters)
  - [core](#core)
  - [ports](#ports)
- [Example Projects](#example-projects)

## Hexagonal Architecture

The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file

> #### [Architecture Docs](docs/architecture.md)

## Project Structure

The following diagrams the project structure

```
.
├── cmd
└── internal
    ├── adapters
    │   ├── primary
    │   └── secondary
    ├── core
    │   ├── domain
    │   ├── entity
    │   ├── service
    │   └── valobj
    └── ports
```

### cmd

The cmd directory contains your applications main modules. The `main.go` file will be responsible for setting up dependencies and running the application

### adapters

The adapters directory contains the **primary** and **secondary** modules that handle communication between external entities and the core of the application. Adapters will be responsible for any data tranformations required for communication, as well as error handling, telemetry, and logging

### core

The core directory contains the applications business logic and will utilize a lite version of Domain Driven Design

There are four main concepts that define the core: **Domains**, **Entities**, **Services**, and **Value Objects**

- `Domains` can be thought of as "orchestrators" for domain use cases. They orchestrate services to implement business rules and validation logic specific to a domain
- `Entities` are the domain objects
- `Services` perform tasks specific to an entity and handle multiple data sources
- `Value Objects` represent shared immutable data types

### ports

The ports directory contains the port definitions that define how the core will interact with internal and external entities. It is split up into the following files:

- `core.go`
  - Contains interfaces that define how services and domains can interact with each other
- `primary.go`
  - Contains interfaces that define how the core allows external entities to interact with it. Ex. `APIPort` defines how the core will allow APIs to interact with it
- `secondary.go`
  - Contains interfaces that define how the core wants to drive/interact with external entities. These entities are usually databases or some other data source, but can also be other internal applications or even a library

## Example Projects

### GoMux

A command-line tool for Tmux setup

> #### [github.com/nxdir-s/gomux](https://github.com/nxdir-s/gomux)

### IdleRpg

An example using a monorepo that contains a websocket game server and a http server for the UI (WIP)

> #### [github.com/nxdir-s/IdleRpg](https://github.com/nxdir-s/IdleRpg)
