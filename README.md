# Go HexArch

This repository serves as a template and documents how to implement Hexagonal Architecture within a golang application

## Hexagonal Architecture

The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file.

## Project Structure

### Adapters

This directory contains the **primary** and **secondary** adapters that handle communication between external entities and the core of the application

### Core

This directory contains the applications core business logic, with the following folder structure

There are four main concepts that define the business logic: Domains, Entities, Services, and Value Objects

- `Domains` can be thought of as "Orchestrators/Interactors" for domain use cases. They orchestrate multiple services and implement business rules and validation logic specific to a domain
- `Entities` are the domain objects
- `Services` handle any actions specific to an entity.
- `Value Objects`

### Ports

This directory contains the port definitions that define how the core will interact with internal and external entities. It is split up into the following files:

- `core.go`
  - Contains interfaces that define how services and domains can interact with each other
- `primary.go`
  - Contains interfaces that define how the core allows external entities to interact with it. Ex. `APIPort` defines how the core will allow APIs to interact with it
- `secondary.go`
  - Contains interfaces that define how the core wants to drive/interact with external entities. These entities are usually databases or downstream data sources, but can also be other internal applications or even a library
