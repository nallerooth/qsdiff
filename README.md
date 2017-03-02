# Query String Diff Tool
A very simple diff tool for comparing query strings.

## Usage

```go
// Results are stored in a `KeyValue` struct
type KeyValue struct {
	key, left, right string
}

// Running Diff will generate a slice of key => KeyValue struct
qsdiff.Diff(strA, strB) => map[string]*KeyValue
```

Example:
```go
strA := "?a=1&b=2"
strB := "?a=1&b=3"

diff := qsdiff.Diff(strA, strB)

// diff contains
diff[a] = *KeyValue{key: "a", left: "1", right: "2"}
diff[b] = *KeyValue{key: "b", left: "1", right: "3"}
```

## Printing
The tool also supports printing the result to the terminal. Values will be 
printed in green or red, depending on whether they match or not.
```go
for _, v := range qsdiff.Diff(strA, strB) {
    v.Print(ignoreMatching) // bool: whether to print matching values
}
```
