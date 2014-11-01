socketpair
==========
This is a simple package for Go that provides an interface to socketpair(2).

Usage
--------
```go
import "github.com/prep/socketpair"
```

```go
func testSocketPair() error {
    sock1, sock2, err := socketpair.New("unix")
    if err != nil {
        return err
    }

    defer sock1.Close()
    defer sock2.Close()

    if _, err := sock1.Write([]byte("Hello World")); err != nil {
        return err
    }

    data := make([]byte, 11)
    if _, err := sock2.Read(data); err != nil {
        return err
    }

    return nil
}
```
