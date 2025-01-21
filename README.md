# Mike's Stax Refund Service

### Summary

This is a Refund Service that lists Payments and Refund Statuses and accepts Refund Requests on behalf of a customer.

This project is simply for demonstration, and some parts are left undone, but documented, for brevity. This service normally wouldn't have need for a database, and would normally map requests to services like a 1st or 3rd Party Authenticator, and a Banking Service, I just added a SQLite Database to simulate the persistance of a downstream service.

### Installation and Running the Project

To setup the `.env` file, run this, but only once!

```bash
cp .env.example .env
```

I've included scripts for fetching the dependencies for, and then running the client, as well as one for the server.

In two terminals, simply run one script each:

```bash
./.dev/client.sh
```

```bash
./.dev/server.sh
```

The client will be on `http://localhost:5173` (the Vite default port), and server on `http://localhost:8080`.

I've used SQLite for simplicity, and the files are in the `db/files` directory if you need to inspect them.


### Points of Interest

The [`RequestService`](internal/sqlite/refund-service.go), while being the focus of the project, is the most intricate, specifically the `RefundPayment` method.

The [`app`](app/services.go) package holds my interfaces, custom errors, and structs, and is implemented both using a live, [SQLite powered implementation](internal/sqlite/payment-client.go), as well as a [Mock implementation](internal/mock/payment-client.go).

I've got [unit tests](internal/sqlite/refund-service_test.go) for the `RefundService.RefundPayment()` method, covering the varying conditions of the code.

Don't look too closely, as it was a means to an end, not the main focus, but I've got a React Client in the `/client` directory, complete with Vite, TailwindCSS, and an API Client library.

### Implementation

The service is a Full-Stack solution, with a REST API in the Gin Framework, SQLite as the database, and a React Client for both simulating the users experience, and hosting the Swagger UI.

I've structured the project after the "Ben Johnson" model, with an `app` package that maintains global struct and service interfaces, and an `internal` package that implements them, divided by their dependencies (`http`, `sqlite`, `mock`).

I used `sqlc` as my "Database Layer", as it allows you to write your queries in raw SQL, but generates the Go code that you would need to build for serialization and field mapping.

### Contribution

`sqlc` is the only dependency outside of Go, and the [installation instructions can be found here](https://docs.sqlc.dev/en/stable/overview/install.html)

To create a query, update the `query file` at [./db/query.sql](./db/query.sql), using `sqlc`'s basic notation, and then run the command 

```bash
sqlc generate
```

to update the `internal/sqlc` package. Note that the code in `internal/sqlc` is generated, and should not be directly updated.

### Security Features

As this is only a demo, the security is more conceptual than actual, but I still built in some safeguards to demonstrate what that might look like. 

1. **Authentication**  
I have an authentication endpoint for requesting, and authentication middleware for verifying user identity, which in our case is a single user with a stubbed AuthToken.

1. **Authorization**  
In the API and the Client, you can only see, and by extention, request refunds for, payments that are attached to your UserID, but I've got verification in the API to ensure that.

1. **Validation**  
You can only request a refund for a payment a single time. Subsequent requests are rejected, with a proper error code.

1. **You can only refund a *Payment*, not an *Amount***  
This seems small, but the Refund Endpoint doesn't accept an amount, rather an ID to an existing payment. That keeps this a "Refund Service" not a "Withdrawl Service".

1. **Refunds are Asynchronous**  
Rather than issuing a refund for every and any payment on the spot, a user the immediate response is a "pending status", as it's reviewed internally. 