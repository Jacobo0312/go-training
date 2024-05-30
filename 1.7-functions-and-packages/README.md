# 1.7 Functions and Packages

## Create a Package

Create a new directory in the `$GOPATH/src` directory called `calculator`. Create a file called `sum.go`. The tree directory should look like this:

```bash
src/
  calculator/
    sum.go
```

Initialize the `sum.go` file with the name of the package:

```go
package calculator
```

You can now start to write the functions and variables for the package. Unlike other programming languages, Go doesn't provide the public or private keywords to indicate if a variable or function can be called from outside or inside of the package. But Go follows two simple rules:

- If you want something to be private, start its name with a lowercase letter.
- If you want something to be public, start its name with an uppercase letter.

So let's add the following code to the `calculator` package we're creating:

```go
package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
    return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
    return number1 + number2
}
```

Let's look at a few things in that code:

- The `logMessage` variable can be called only from within the package.
- The `Version` variable can be reached from anywhere. We recommend that you include a comment to describe the purpose of this variable. (This description is useful for anyone who uses the package.)
- The `internalSum` function can be called only from within the package.
- The `Sum` function can be reached from anywhere. We recommend that you include a comment to describe the purpose of the function.

To confirm that everything is working, you can run the `go build` command in the `calculator` directory. If you do, notice that no executable binary is generated.

## Create a Module

You've placed the calculator functionality inside a package. Now it's time to put that package into a module. Go modules typically contain packages that offer related functionality. A package's module also specifies the context that Go needs to run the code you've grouped together. This contextual information includes the Go version your code is written for.

Also, modules help other developers reference specific versions of your code and make working with dependencies easier. Another benefit is that our program's source code doesn't strictly need to exist in the `$GOPATH/src` directory. Freeing that restriction makes it more convenient to work with different package versions in other projects at the same time.

So, to create a module for the `calculator` package, run this command in the root directory (`$GOPATH/src/calculator`):

```bash
go mod init github.com/myuser/calculator
```

After you run this command, `github.com/myuser/calculator` becomes the module's name. You'll use that name to reference it in other programs. The command also creates a new file called `go.mod`. Finally, the tree directory now looks like this:

```bash
src/
  calculator/
    go.mod
    sum.go
```

The contents of the `go.mod` file should look like the following code. (The Go version might be different.)

```go
module github.com/myuser/calculator

go 1.14
```

To reference your `calculator` package in other programs, you need to import it by using the module name. In this case, the name is `github.com/myuser/calculator`. Now let's look at an example of how to use this package.

## Reference a Local Package (a Module)

Now let's use the package. We'll continue with the sample application we've been using. This time, instead of having the `sum` function in the main package, let's use the one we created earlier in the `calculator` package.

The tree file structure now should look like this:

```bash
src/
  calculator/
    go.mod
    sum.go
  helloworld/
    main.go
```

We'll use this code for the `$GOPATH/src/helloworld/main.go` file:

```go
package main

import (
  "fmt"
  "github.com/myuser/calculator"
)

func main() {
    total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)
}
```

Notice that the import statement uses the name of the package you created: `calculator`. To call the `Sum` function from that package, you need to include the package name, as in `calculator.Sum`. Finally, you now also have access to the `Version` variable. You call it like this: `calculator.Version`.

If you try to run the program now, it won't work. You need to tell Go that you're using modules to reference other packages. To do so, run this command in the `$GOPATH/src/helloworld` directory:

```bash
go mod init helloworld
```

In the previous command, `helloworld` is the name of the project. This command creates a new `go.mod` file, so now the tree directory looks like this:

```bash
src/
  calculator/
    go.mod
    sum.go
  helloworld/
    go.mod
    main.go
```

When you open the `go.mod` file, you should see something like the following code. (The Go version might be different.)

```go
module helloworld

go 1.14
```

Because you're referencing a local copy of the module, you need to inform Go that you don't want to use a remote location. So you need to manually modify the `go.mod` file to include the reference, like this:

```go
module helloworld

go 1.14

require github.com/myuser/calculator v0.0.0

replace github.com/myuser/calculator => ../calculator
```

The `replace` keyword specifies to use a local directory instead of a remote location for the module. In this case, because the `helloworld` and `calculator` programs are in `$GOPATH/src`, the location is simply `../calculator`. If the module's source is in a different location, you define the local path here.

Run the program by using this command:

```bash
go run main.go
```

The output should be as follows:

```bash
8
Version:  1.0
```
