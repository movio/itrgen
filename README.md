# itrgen - Iterator Generator

Typesafe code generators for common slice fuctions such as .Map(fn)

## Installation

```shell
go get -u github.com/movio/itrgen/.../
```

## Usage

Add go:generate commands to a file eg,

```go
package main
type MyStruct struct {
    Name string
}
//go:generate itrgen -type MyStruct -map-to="MyStruct,string,int64"
```

then

```shell
go generate
```

then

```go
// Usage
// ...
list := MyStructItr{MyStruct{"John"}, MyStruct{"Jane"}}
hello, _ := list.MapString(func(s MyStruct) (string, error) {
   return fmt.Sprintf("hello %v", s.Name)
})
// ...
```
    
### Arguments

    Usage of itrgen:
      itrgen [-package <pkg>] -type <type> [-map-to="<one>,<two>"]
      -package string [Optional]
            the package that this Itr will be in. (default "main")
      -type string
            the main type that we want to map *from*. eg, string or MyStruct
      -map-to string
            comma separated list of types that can be mapped *to*



## Exmaples

```go
//go:generate itrgen -package="main" -type="string" -map-to="MyStruct,string,int64"
// Will create
type StringItr []string
func (itr StringItr) MapMyStruct(fn func(string) (MyStruct, error)) ([]MyStruct, error) { ... }
func (itr StringItr) MapString(fn func(string) (string, error)) ([]string, error) { ... }
func (itr StringItr) MapInt64(fn func(string) (int64, error)) ([]int64, error) { ... }

//go:generate itrgen -type=MyStruct -map-to="MyStruct,string,int64"
// Will create
type MyStructItr []MyStruct
func (itr MyStructItr) MapMyStruct(fn func(MyStruct) (MyStruct, error)) ([]MyStruct, error) { ... }
func (itr MyStructItr) MapString(fn func(MyStruct) (string, error)) ([]string, error) { ... }
func (itr MyStructItr) MapInt64(fn func(MyStruct) (int64, error)) ([]int64, error) { ... }
```
    
### Run the Examples

```shell
cd examples
rm *_itr.go
go generate
go test
```
