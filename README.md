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

