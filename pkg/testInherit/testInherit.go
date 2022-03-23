package main

import (
    "fmt"
)

type Dad interface {
    GetName() string
    Print()
}

type DadImpl struct{}
func (DadImpl) GetName() string { return "dad" }
func (d DadImpl) Print()        { fmt.Println("print name:", d.GetName()) }

type SonImpl struct{DadImpl}
func (s *SonImpl) GetName() string { return "son" }

//func (s *SonImpl) print()        { fmt.Println("name:", s.getName()) }

var _ Dad = (*SonImpl)(nil)

func main() {
    fmt.Println("begin==================")
    var dad Dad = &SonImpl{}
    fmt.Println("dad name:", dad.GetName())
    dad.Print()

}
