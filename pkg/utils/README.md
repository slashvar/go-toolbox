# utils

Generic helpers for slices and optional values.

```go
import "github.com/slashvar/go-toolbox/pkg/utils"
```

## Slices

Functions taking a predicate or accumulator receive it first, before the slice.

| Function                 | Description                                                                                |
|--------------------------|--------------------------------------------------------------------------------------------|
| `Map(in, f)`             | Returns a new slice of `f(e)` for each `e` in `in`                                         |
| `Iter(s, f)`             | Calls `f` on each element                                                                  |
| `Accumulate(f, init, s)` | Left fold: `f(...f(f(init, s[0]), s[1])..., s[n])`                                         |
| `ForAll(f, s)`           | Reports whether `f` holds for every element                                                |
| `Exists(f, s)`           | Reports whether `f` holds for at least one element                                         |
| `Filter(f, s)`           | Returns the elements for which `f` holds                                                   |
| `OptionalFilter(f, s)`   | Returns the values of the non-empty options `f(e)`                                         |
| `Reverse(s)`             | Reverses `s` in place                                                                      |
| `Insert(x, pos, s)`      | Inserts `x` at `pos`, appends when `pos >= len(s)`, panics when `pos < 0`                  |
| `LowerBound(x, s)`       | Binary search in a sorted `s`, returns the index of the first element not smaller than `x` |
| `Partition(x, s)`        | Moves elements smaller than `x` first and returns the split index                          |
| `PartitionFilter(f, s)`  | Moves elements satisfying `f` first and returns the split index                            |
| `Combine(s1, s2, zip)`   | Returns `zip(s1[i], s2[i])` over the shorter length                                        |

Notes:

- `Insert`, `Reverse`, `Partition` and `PartitionFilter` modify the slice in place. `Partition` and `PartitionFilter` do not preserve order.
- `LowerBound` returns an undefined result on an unsorted slice. With duplicates, it returns the first matching index.
- `LowerBound` and `Partition` accept any `cmp.Ordered` type.

## Options

`Option[T]` models a value that may be absent. `NewOption(v)` holds a value, `NilOption[T]()` is empty.

| Method or function      | Description                                               |
|-------------------------|-----------------------------------------------------------|
| `HasValue()`            | Reports whether the option holds a value                  |
| `Value()`               | Returns the value, panics when the option is empty        |
| `Do(f)`                 | Calls `f(Value())` if non-empty, returns the option       |
| `Else(f)`               | Calls `f()` if empty, returns the option                  |
| `OptionalMap(o, f)`     | Returns an option holding `f(o.Value())`, or an empty one |
| `OptionalFlatMap(o, f)` | Same as `OptionalMap` for `f` returning an `Option`       |
| `OptionalDo(o, f)`      | Function form of `o.Do(f)`                                |
| `OptionalElse(o, f)`    | Function form of `o.Else(f)`                              |

`Do` and `Else` chain:

```go
utils.NewOption(42).
	Do(func(v int) { fmt.Println("got", v) }).
	Else(func() { fmt.Println("empty") })
```

An empty option is a nil `*OptionImplem[T]`. Stored in an `Option[T]` interface, it compares non-nil: test it with `HasValue()`, never `== nil`.
