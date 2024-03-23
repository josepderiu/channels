# Learning Channels in Go

This project is a learning resource for understanding how to use channels in Go. Channels are a powerful feature in Go that allow goroutines to communicate with each other and synchronize their execution.

## Project Structure

The project consists of two main files:

- `main.go`: This file contains the implementation of two worker functions, `workerOne` and `workerTwo`, which send messages to their respective channels at different intervals. The `main` function starts these workers as goroutines and listens for their messages. After a certain period of time, the program stops listening and exits.

- `main_test.go`: This file contains tests for the worker functions. The tests start the workers as goroutines and check that they send the expected messages within a certain timeout.

## Running the Project

To run the project, use the `go run` command:

go run main.go

This will start the program and you will see messages from the worker functions printed to the console.

## Running the Tests

To run the tests, use the `go test` command:

go test [-v] [--cover]

This will run the tests and print the results to the console.

## Learning Resources

For more information on channels in Go, check out the following resources:

- [Go by Example: Channels](https://gobyexample.com/channels)
- [A Tour of Go: Concurrency](https://tour.golang.org/concurrency/2)
- [Go Documentation: Channel Types](https://golang.org/ref/spec#Channel_types)