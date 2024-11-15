# Test utils

Common test functions and variables.

```
go get -u github.com/a-novel-kit/test-utils
```

- [Run in external executable](#run-in-external-executable)
- [GRPC utils](#grpc-utils)
  - [Check GRPC status](#check-grpc-status)
  - [Wait for GRPC service readiness](#wait-for-grpc-service-readiness)

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/a-novel-kit/test-utils/main.yaml)
[![codecov](https://codecov.io/gh/a-novel-kit/test-utils/graph/badge.svg?token=ZfQNEOQcW8)](https://codecov.io/gh/a-novel-kit/test-utils)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/a-novel-kit/test-utils)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/a-novel-kit/test-utils)

![Coverage graph](https://codecov.io/gh/a-novel-kit/test-utils/graphs/sunburst.svg?token=ZfQNEOQcW8)

# Run in external executable

Sometimes, it is useful to test a function in an isolated environment. `RunCMD` will run the
current test in a separate executable, with a controlled environment.

```go
package my_test

import (
  "testing"

  testutils "github.com/a-novel-kit/test-utils"
)

func TestSomething(t *testing.T) {
  testutils.RunCMD(t, &testutils.CMDConfig{
    CmdFn: func(t 8testing.T) {
      // Runs under the separate executable.
    },
    MainFn: func(t *testing.T, res *testutils.CMDResult) {
      // More information available in the documentation of testutils.CMDResult.
    }, 
    // Optional, custom environment.
    Env: []string{"FOO=BAR"},
  })
}
```

> Note: in the above example, `TestSomething` will be triggered twice. It is important that the 
> `RunCMD` util is only called once per test function.

# GRPC utils

Helpers for testing GRPC services.

## Check GRPC status

Check the error returned by a GRPC call, based on a target status. If the status targeted is
`codes.OK`, the error must be nil.

```go
package my_test

import (
  "testing"

  "google.golang.org/grpc/codes"

  testutils "github.com/a-novel-kit/test-utils"
)

func TestSomething(t *testing.T) {
  _, grpcErr := makeGRPCCall()
  
  // Ensure the call returned without an error.
  testutils.RequireGRPCCodesEqual(t, grpcErr, codes.OK)
}
```

## Wait for GRPC service readiness

When running a local server for testing, this utils waits for the server to be ready before
running the tests.

```go
package my_test

import (
  "testing"

  "google.golang.org/grpc"

  testutils "github.com/a-novel-kit/test-utils"
)

func TestSomething(t *testing.T) {
  var conn *grpc.ClientConn
  
  // Init the conn...
  
  // Mark the test as error if server fails to report healthy after a certain time.
  testutils.WaitConn(t, conn)
}
```
