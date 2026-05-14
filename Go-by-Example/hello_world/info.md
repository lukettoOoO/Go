# Quick Reference: Getting Started with Go

### 1. Setup & Environment
* **Prerequisites:** A text editor (VSCode, GoLand, or Vim) and a command terminal.
* **Installation:** Install the latest version via the official [go.dev](https://go.dev) installer.

### 2. Initializing a Project
All Go code belongs to a module, which manages dependencies.

1.  **Create a workspace:**
    ```bash
    mkdir hello
    cd hello
    ```
2.  **Initialize the module:**
    Run the `go mod init` command followed by your module path (typically your repository location, e.g., `github.com/user/project`).
    ```bash
    go mod init example/hello
    ```
    * *Result:* Creates a `go.mod` file to track versions and dependencies.

### 3. Writing "Hello, World"
Create a file named **`hello.go`**:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}