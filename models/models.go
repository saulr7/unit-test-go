package models

type Dog struct {
	Name string
	Age  int
	Kind Kind
}

type Kind struct {
	Name string
}
