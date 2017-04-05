package main

type MyStruct struct {
	Name string
}

//go:generate itrgen -type=MyStruct -map-to="MyStruct,string,int64"
//go:generate itrgen -package="main" -type="string" -map-to="MyStruct,string,int64"

func main() {}
