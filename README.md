# Template Funcs

A collection of useful template functions for Go’s `html/template` package.

Unlike other such libraries, this is split into multiple modules so you can pick and choose which ones to add.
Selectively importing only the modules/functions you need can help reduce the number of dependencies & resulting binary size.

## Get Started

### Installation

```bash
# Install only the strings module
go get go.chimbori.app/templatefuncs/string
```

### Usage

Import the desired modules and register their function maps with your templates:

```go
import (
    "html/template"
    stringfuncs "go.chimbori.app/templatefuncs/string"
)

tmpl := template.New("example").Funcs(stringfuncs.FuncMap())
```

## Available Modules

### `string`

String manipulation functions for common template operations.

`go get go.chimbori.app/templatefuncs/string`

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

## Contributing

Contributions are welcome! Please ensure:
- Ensure that all tests continue to pass: `task test`. Install [Task](https://taskfile.dev/) if not already available.
- Please add comprehensive tests & documentation for any newly-added functions.
