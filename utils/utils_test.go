package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"
	s := struct{ Testing string }{Testing: "test"}

	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)
		if x != hash {
			t.Errorf("Expected  %s, got %s", hash, x)
		}
	})

	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

// ExampleHash is generated on godocs file as example
func ExampleHash() {
	s := struct{ Testing string }{Testing: "test"}
	x := Hash(s)
	fmt.Println(x)

	// Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("Tobytes should return a slice of bytes, but got %s", k)
	}
}

func TestSpliter(t *testing.T) {
	type test struct {
		input  string
		sep    string
		index  int
		output string
	}
	tests := []test{
		{input: "0:6:0", sep: ":", index: 1, output: "6"},
		{input: "0:6:0", sep: ":", index: 10, output: ""},
		{input: "0:6:0", sep: "/", index: 0, output: "0:6:0"},
	}
	// table test -> table test는 여러 테스트 케이스를 한번에 for loop를 통해서 테스트 하는 것을 말한다.
	for _, tc := range tests {
		result := Splitter(tc.input, tc.sep, tc.index)
		if result != tc.output {
			t.Errorf("Expected %s, but got %s", tc.output, result)
		}
	}
}

func TestHandleError(t *testing.T) {
	oldLogFn := logFn
	defer func() {
		logFn = oldLogFn
	}()
	called := false
	logFn = func(v ...interface{}) {
		called = true
	}
	err := errors.New("test")
	HandleError(err)
	if !called {
		t.Errorf("HandleError should call fn")
	}
}
