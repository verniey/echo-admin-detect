# Echo Middleware: Admin Check and Days Left

This middleware for the Echo framework in Go provides two functionalities:

1. Calculate the number of days left until 1 Jan 2025 and return it in the response.
2. Checks if the HTTP header `User-Role` contains "admin". If it does, it logs "Red button user detected".


## Requirements

- Go installed on your machine
- Echo framework installed (`go get -u github.com/labstack/echo/v4`)

## Installation

1. Clone this repository:


2. Navigate to the directory:

cd <project_directory>/cmd/mw-task

where main.go file is located


3. Build and run the application:

go build -o <executable_name>
./<executable_name>

Alternatively, if you don't want to build the executable and then run it, you can use:

`go run main.go`

## Usage

### 1. Calculate Days Left Handler

To get the number of days left until 1 Jan 2025, make a GET request to the endpoint `/days`.

Example:

GET /days


Response:

Days left until 1 Jan 2025: <number_of_days>