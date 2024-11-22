## Hexagonal Architecture

The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file.

The pattern allows us to isolate the core logic of our application from outside concerns. Having our core logic isolated means we can easily change data source details **without a significant impact or major code rewrites to the codebase**.

One of the main advantages in having an app with clear boundaries is implementing a testing strategy — the majority of our tests can verify our business logic **without relying on protocols that can easily change**.

## Intent

Allow an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases.

When any driver wants to use the application at a port, it sends a request that is converted by an adapter for the specific technology of the driver into an usable procedure call or message, which passes that to the application port. The application is blissfully ignorant of the driver’s technology. When the application has something to send out, it sends it out through a port to an adapter, which creates the appropriate signals needed by the receiving technology (human or automated). The application has a semantically sound interaction with the adapters on all sides of it, without actually knowing the nature of the things on the other side of the adapters.

## Ports

The term “ports” simply refers to entry points to the application core. They contain (typically technology-neutral) interfaces that make it possible for external entities to obtain a set of rules for communicating with the core. Since the ports are essentially just gateways, another agent is necessary to actually make the communication happen. These are adapters.

## Adapters

The adapters actively initiate the communication between external entities and the core. Each port can serve many adapters. A common example of a controller could be a REST controller or any other API request handler.

Here is an extremely important thing about the adapters architecture – ports/adapters work with both the external systems that start the communication (driving side) and the ones that receive it (driven side). But the exact mechanism slightly differs.

## Primary/Driving side

The driving actors are those that **start the interaction** with the application by initiating a query. For example, it can be the mobile application interface or user interface code of a web app. The user input passed into the UI is taken by the adapter and sent to the core through the port. Both the port (interface) and the implementation of the interface will be inside the core/hexagon.

## Secondary/Driven side

The driven actors are those that need the core application to interact with them. It could be databases or even other applications. The application calls the external (driven) entity, and the adapter implements the port for the core to use. The implementation is within the driven adapter.

## Benefits

- When done correctly, it makes it possible to isolate the application and business logic from external factors so that they can all be **tested** easily and separately.
- At the same, their dependencies can be easily **mocked**.
- Designing the user interfaces by their purpose rather than technology ensures that your application’s technology stack can freely grow over time.
- Helps implement **Domain-Driven Design** by making sure that the domain logic does not leak out of the core.
- The ports and adapters are just as replaceable as all the external entities, further contributing to the **scalability** of the entire application.
- The advanced separation of concerns also makes the app **easier to maintain**, as changing the code in one place or adding new dependencies/ways to interact with the app, do not require significant code changes.
- Since one can test outside dependencies without any extra mocking tools, it improves the overall **testability** of the application.
