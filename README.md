# hello_world

This is my first Go application.

## Getting started

To run programs the source code and dependencies need to be compiled to executable binary.

The following goes in the command line:

```commandline
  go build <file name>
  ./<file name>
```

The **'go build'** command will cmpile for us. To execute type in **'./'** followed by the binary file name.

## Variables

You can declare variables with the following:

```commandline

var a int

var (
	b bool
	c float32
	d string
)

```

Name the varialbe and define its data type.

Here is a sample of a the shorthand operator **':='** which assigns and initialize in one line.

```commandline

  a := 24
  b := false
  c := 23.4
  d := "this is a diffrent string"

```

