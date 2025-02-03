# Integration Test in Golang Using Ginkgo and Memongo

This repository demonstrates an **integration test** in Go using **Ginkgo** for behavior-driven development (BDD) testing and **Memongo** for running an **in-memory MongoDB** instance.

### Dependencies Used

- **Ginkgo**: [https://github.com/onsi/ginkgo/v2](https://github.com/onsi/ginkgo/v2)
- **Memongo**: [https://github.com/tryvium-travels/memongo](https://github.com/tryvium-travels/memongo)

### Running the Tests

Once the dependencies are installed, you can run the tests using the following command:

```bash
go test ./...
```

This will execute all tests in the project, including the Ginkgo-based tests with the in-memory MongoDB provided by Memongo.

