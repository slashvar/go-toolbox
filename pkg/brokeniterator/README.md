# brokeniterator

Adapter from `Next`-style iterators to `iter.Seq`.

```go
import "github.com/slashvar/go-toolbox/pkg/brokeniterator"
```

## API

| Name                | Description                                                                   |
|---------------------|-------------------------------------------------------------------------------|
| `BrokenIterator[T]` | Interface with `Next() (T, bool)`: the current value, and whether more follow |
| `Seq(b)`            | Returns an `iter.Seq[T]` over `b`, usable in a `for range` loop               |

`Seq` calls `Next` once per element and stops calling it when the loop exits early.

`Seq` always yields the first `Next` result. An empty iterator yields one zero value.

## Example

`runtime.Frames` follows the `Next` idiom:

```go
pcs := make([]uintptr, 16)
n := runtime.Callers(0, pcs)
for f := range brokeniterator.Seq[runtime.Frame](runtime.CallersFrames(pcs[:n])) {
	fmt.Println(f.Function)
}
```
