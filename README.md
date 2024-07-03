# Smartcloud Prices API

This project implements an API for fetching and returning machine instance prices from Smartcloud using the Gin framework.

## Assumptions
- The Smartcloud API endpoint is `https://api.smartcloud.com`.
- The `kind` parameter is required to fetch prices for specific machine instances.

## Design Decisions
- Used Gin framework for its simplicity and performance.
- Handled errors gracefully by returning appropriate HTTP status codes.

## Running the Code
1. Clone the repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to install dependencies.
4. Start the server with `go run main.go`.
5. Access the API at `http://localhost:8080/prices?kind=<instance-kind>`.

## Example
```sh
curl http://localhost:8080/prices?kind=sc2-micro
