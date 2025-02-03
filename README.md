# Integration Test in Golang Using Ginkgo and Memongo

This repository demonstrates an **integration test** in Go using **Ginkgo** for behavior-driven development (BDD) testing and **Memongo** for running an **in-memory MongoDB** instance.

### Installation

Install the dependencies using `go get`:

   ```bash
   go get github.com/onsi/ginkgo/v2
   go get github.com/tryvium-travels/memongo
   ```

This will install **Ginkgo** for writing the tests and **Memongo** for running the in-memory MongoDB instance.

### Running the Tests

Once the dependencies are installed, you can run the tests using the following command:

```bash
go test ./...
```

This will execute all tests in the project, including the Ginkgo-based tests with the in-memory MongoDB provided by Memongo.