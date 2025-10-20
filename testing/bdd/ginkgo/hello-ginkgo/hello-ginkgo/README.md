# Hello Ginkgo

This project demonstrates a simple example of using Ginkgo for Behavior-Driven Development (BDD) in Go. It includes a basic application that provides a greeting function and a test suite that verifies its behavior.

## Project Structure

```
hello-ginkgo
├── src
│   └── main.go          # Contains the main application code
├── tests
│   └── hello_ginkgo_test.go  # Contains the Ginkgo test suite
├── go.mod               # Module definition for the Go project
└── README.md            # Project documentation
```

## Getting Started

To get started with this project, ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Running the Application

To run the application, navigate to the `src` directory and execute the following command:

```
go run main.go
```

### Running the Tests

To run the tests using Ginkgo, navigate to the `tests` directory and execute:

```
go test
```

Make sure you have Ginkgo installed. If you haven't installed it yet, you can do so by running:

```
go get -u github.com/onsi/ginkgo/ginkgo
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.