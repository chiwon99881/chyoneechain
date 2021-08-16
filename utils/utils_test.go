package utils

import (
	"encoding/hex"
	"encoding/json"
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

func TestFromBytes(t *testing.T) {
	type testStruct struct {
		Test string
	}
	var restored testStruct
	ts := testStruct{"test"}
	b := ToBytes(ts)
	FromBytes(&restored, b)
	// DeepEqual은 두 interface가 같은지 확인한다 interface의 필드와 타입 unexport, export까지 모두 다
	if !reflect.DeepEqual(ts, restored) {
		t.Error("FromBytes() should restore struct")
	}
}

func TestToJSON(t *testing.T) {
	type testStruct struct {
		Test string
	}
	s := testStruct{"test"}
	b := ToJSON(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("Expected %v, but got %v", reflect.Slice, k)
	}
	var restored testStruct
	json.Unmarshal(b, &restored)
	if !reflect.DeepEqual(restored, s) {
		t.Error("ToJSON() should encode to JSON correctly.")
	}
}
