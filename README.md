# nushi-gh-test

A throwaway repository for testing [Nushi](https://github.com/afrase/nushi)'s GitHub integration: pull requests, issues, review threads and checks.

Two checks run on every push and pull request:

- **test** runs `go test ./...` and passes.
- **on-demand** fails when the branch has a file named `FAIL`, and passes otherwise.
