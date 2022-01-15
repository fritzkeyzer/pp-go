# A Pretty Printer for Golang

Uses github.com/fatih/color for coloring on various terminals

## Make ur dumps pretty:

```go

type someDumbStruct struct{
    strVal string
    intVal int
    flVal float64
}

s := someDumbStruct{
    strVal: "blah blah blah",
    intVal: 1337,
    flVal:  42.0,
}

pp.Dump(&s)
```

![Screenshot](https://github.com/fritzkeyzer/pp-go/blob/master/screenshot.png)
