# Testing

- Test must start with word Test and the next word must be capital letter. Ex: Testget... doesn't work
- In go there is no asserts because it stops the rest of the test if one fails
- We should have 1 test case per return
## Test cases
In any programming language you have 3 different steps that you must comply:
- Initialization
- Execution
- Validation
## go test
Runs a test method.
## go test -cover
Runs a test with coverage information.
## Libraries
### stretchr/testify/assert
For assert test.

# Benchmark
## go test -bench=.