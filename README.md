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

Name the variable and define its data type.

Here is a sample of a the shorthand operator **':='** which assigns and initialize in one line.

```commandline

  a := 24
  b := false
  c := 23.4
  d := "this is a diffrent string"

```

## Data Types

Examples on how to declare different datatypes

```

// user specified types
 const a int32 = 12        // 32 bit
 const b float32 = 34.2    // 32 bit float
 var c complex128 = 1 + 4i // 128 complex number
 var d uint16 = 14         // 16 bit unsigned integer

 // default types
 n := 42                    // int
 pi := 3.14                 // float54
 x, y := true, false        // bool
 z := "this is a Go string" // string

 fmt.Printf("user-specified \n %T %T %T %T\n", a, b, c, d)
 fmt.Printf("default \n %T %T %T %T %T\n", n, pi, x, y, z)

```

Note the **'%T'** syntax in the Printf. This is called a verb and it will print the data type of the corresponding variable that is printed.

## Arrays

Arrays are fixed sized meaning they cannot grow of shrink.

```commandline

//defining an array that stores 4 items

  var testArr = [4]string{"yes", "no", "stop", "naur"}

  // loop through
  fmt.Println("for loop")
  for i := 0; i < len(testArr); i++ {
  item := testArr[i]
  fmt.Println(i, item)
 }

 var testArr1 = [...]string{"what", "life", "donut", "test"}

 fmt.Println("range")

 for index, item := range testArr1 {
  fmt.Println(index, item)
 }

```