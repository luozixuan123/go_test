package test

import (
	"testing"
	"fmt"
	)

func add (a int ,b int) (int){
	return a + b
}

func TestAdd(t *testing.T) {
	if 7 != add(3, 3){
		t.Error("add fail")
	}else{
		fmt.Println("add success")
	}
}