# Template Funcs

A collection of useful template functions for Go’s `html/template` package.

Unlike other such libraries (e.g.
[Masterminds/sprig](https://github.com/Masterminds/sprig/blob/master/go.mod)),
this one is split into multiple modules so you can pick and choose the ones to add.

Selectively importing only the modules/functions you need can help reduce the number of dependencies & resulting binary size.

## Installation

```bash
# stdfuncs has no dependencies beyond the stdlib.
go get chimbori.dev/templatefuncs/stdfuncs
```

## Usage

```go
import (
    "html/template"
    "chimbori.dev/templatefuncs/stdfuncs"
)

tmpl := template.New("example").Funcs(stdfuncs.FuncMap())
```

# Available Modules

## stdfuncs

Functions that do not have any external dependencies beyond the Go standard library.

`go get chimbori.dev/templatefuncs/stdfuncs`

- `replace`: Replaces all occurrences of a substring with another string.

  ```go
  {{ "hello world" | replace "world" "Go" }}
  // Output: "hello Go"

  {{ "banana" | replace "a" "o" }}
  // Output: "bonono"

  {{ "hello.world.test" | replace "." "!" }}
  // Output: "hello!world!test"
  ```

- `cat`: Concatenates multiple values into a single space-separated string. Automatically filters out nil values.

  ```go
  {{ cat "hello" "world" "from" "Go" }}
  // Output: "hello world from Go"

  {{ cat "foo" nil "bar" }}
  // Output: "foo bar"

  {{ cat "Status:" .Status }}
  // Output: "Status: active"
  ```

# Contributing

Contributions are welcome!
- Ensure that all tests continue to pass: Run `task test`.
  - Install [Task](https://taskfile.dev/) if not already installed.
- Please add comprehensive tests & documentation for any newly-added functions.
