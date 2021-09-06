# gomap

Map functions utilizing basic interfaces, assertions, and ptypes.

This code and [ptypes](https://github.com/prequist/ptypes) are both artistic bloat,
utilised for DX concepts and/or type-evading for operations.

## Usage

<details closed>
<summary>Basic Setup</summary>

```go
import (
	"github.com/prequist/gomap"
)

func Showcase() {
	// gomap.Transformer
	transformer := func(i interface{}) interface{} {
		// type assert the interface and add 1.
		return i.(int) + 1
    }
    // Create a new list
    list  := gomap.New(1, 2, 3, 4)
    mappable := gomap.MappableList{list}
    // The outcome.
    outcome := mappable.Map(transformer)
    
    // For predefined slices, we can use this flow:
    slice := []int{1, 2, 3, 4, 5}
    list = gomap.New(slice)
    mappable = gomap.MappableList{list}
    outcome := mappable.Map(transformer)
}
```

</details>

<details closed>
<summary>Boxed</summary>

```go
import (
	"github.com/prequist/gomap"
	"github.com/prequist/ptypes"
)

func MakeBoxedAndConvert() []int {
	e := gomap.NewBoxed(1, 2, 3, 4, 5)
	mappable := gomap.MappableList{e}
	mappable.Map(func(v interface{}) interface{} {
		ib := v.(ptypes.Box).IntBox()
		return ptypes.FromInt(*ib.Int() + 2)
	})
	arr := make([]int, len(e.Items()))
	for index, vi := range e.Items() {
		arr[index] = *vi.(ptypes.Box).IntBox().Int()
	}
	return arr
}
```

</details>

<details closed>
<summary>Using <code>interface{}</code>s</summary>

```go
import (
	"github.com/prequist/gomap"
	"github.com/prequist/ptypes"
)

func MakeAndConvert() []int {
	e := gomap.New(1, 2, 3, 4, 5)
	mappable := gomap.MappableList{e}
	mappable.Map(func(v interface{}) interface{} {
		return v.(int) + 1
	})
	arr := make([]int, len(e.Items()))
	for index, vi := range e.Items() {
		arr[index] = vi.(int)
	}
	return arr
}
```

</details>

## Benchmarks

These are probably a bit bloat over time, however, they're not horrendous (compared to other)
implementations.

On Mac OSX (Monterey, M1)

### map_test.go

```
goos: darwin
goarch: arm64
pkg: github.com/prequist/gomap
BenchmarkPlain-8                12474228               100.0 ns/op
BenchmarkBox-8                    759380              1523 ns/op
BenchmarkString-8                4029303               277.8 ns/op
BenchmarkBoxedString-8            782708              1501 ns/op
BenchmarkAdd-8                  17820822                94.33 ns/op

```

### iterator_test.go

```
BenchmarkIteratorNext-8         1000000000               0.5661 ns/op
```


### filter_test.go

```
BenchmarkFilter
BenchmarkFilter/Filter
BenchmarkFilter/Filter-8         	1000000000	         0.0000017 ns/op
```