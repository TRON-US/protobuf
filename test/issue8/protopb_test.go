// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package proto

import (
	fmt "fmt"
	_ "github.com/tron-us/protobuf/gogoproto"
	github_com_tron_us_protobuf_jsonpb "github.com/tron-us/protobuf/jsonpb"
	github_com_tron_us_protobuf_proto "github.com/tron-us/protobuf/proto"
	proto "github.com/tron-us/protobuf/proto"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestFooProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, false)
	dAtA, err := github_com_tron_us_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo{}
	if err := github_com_tron_us_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_tron_us_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestFooJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	marshaler := github_com_tron_us_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo{}
	err = github_com_tron_us_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestFooProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	dAtA := github_com_tron_us_protobuf_proto.MarshalTextString(p)
	msg := &Foo{}
	if err := github_com_tron_us_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFooProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	dAtA := github_com_tron_us_protobuf_proto.CompactTextString(p)
	msg := &Foo{}
	if err := github_com_tron_us_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
