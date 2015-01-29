package bitops

import (
	"fmt"
	"testing"
)

func TestSetBit(t *testing.T) {
	var testInt uint64
	testInt = 32

	SetBit(&testInt, 5, false)
	if testInt != uint64(0) {
		t.Errorf("Failed set bit to false test")
	}
	SetBit(&testInt, 5, false)
	if testInt != uint64(0) {
		t.Errorf("Failed set bit to false again test")
	}

	testInt = 0
	SetBit(&testInt, 7, true)
	if testInt != uint64(128) {
		t.Errorf("Failed set bit to true test")
	}
	SetBit(&testInt, 7, true)
	if testInt != uint64(128) {
		t.Errorf("Failed set bit to true again test")
	}
}

func TestQueryBit(t *testing.T) {
	var testInt uint64
	testInt = 32

	for i := 0; i < 64; i++ {
		fmt.Println(QueryBit(&testInt, i))
	}

	if !QueryBit(&testInt, 2) {
		t.Error("Failed query bit false test")
	}

	if QueryBit(&testInt, 4) {
		t.Error("Failed query bit true test")
	}
}

func TestFlipAll(t *testing.T) {
	var testInt uint64
	testInt = 32

	for i := 0; i < 64; i++ {
		SetBit(&testInt, i, true)
	}
	
	if ^uint64(0) != testInt {
		t.Error("Failed to flip all")
	}
}
