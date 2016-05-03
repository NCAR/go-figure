package gofigure

import (
	"github.com/spf13/viper"
	"sort"
	"testing"
)

func TestGather(t *testing.T) {
	type x struct {
		f, e sort.StringSlice
	}

	tests := map[string]x{
		"ok":     x{f: gather([]string{"_tests/gather"}), e: sort.StringSlice{"_tests/gather/a", "_tests/gather/file", "_tests/gather/that", "_tests/gather/this"}},
		"sorted": x{f: gather([]string{"_tests/yaml", "_tests/json"}), e: sort.StringSlice{"_tests/yaml/1.yml", "_tests/yaml/2.yml", "_tests/json/1.json", "_tests/json/2.json"}},
		"bad":    x{f: gather([]string{"not-a-directory"})},
	}

	for label, x := range tests {
		t.Log("Running checks on %q", label)
		if len(x.f) != len(x.e) {
			t.Errorf("Lengths do not match")
		}
		for i := range x.f {
			if x.f[i] != x.e[i] {
				t.Logf("x.f=%v", x.f)
				t.Logf("x.e=%v", x.e)
				t.Fatalf("Expected %q got %q", x.f[i], x.e[i])
			}
		}
	}
}

func TestParse(t *testing.T) {
	type x struct {
		v map[string]interface{}
		p []string
		f string
	}

	tests := map[string]x{
		"yaml 1": x{
			p: []string{"_tests/yaml"},
			f: "yaml",
			v: map[string]interface{}{
				"joe.hacker":               false,
				"joe.age":                  12,
				"joe.beard":                true,
				"joe.clothing.shirt.color": "red",
				"bob.hacker":               false,
			},
		},
		"yaml 2": x{
			p: []string{"_tests/does-not-exist"},
			f: "yaml",
		},
	}

	testor := func(x x) {

		v := viper.New()
		v.SetConfigType(x.f)
		Parse(v, x.p)
		for key, val := range x.v {
			if va := v.Get(key); va != val {
				t.Errorf("Mis-match.  Got '%v' wanted '%v'", va, val)
			}
		}
	}

	for name, x := range tests {
		t.Log("Running checks %q", name)
		testor(x)
	}
	// type x struct {
	// f, e
	// }
	// if v, e := Parse("yaml", []string{"test-a"}); e == nil {
	// fmt.Println(v.Get("joe.hacker"), e)
	// fmt.Println(v)
	// } else {
	// t.Errorf("e==", e)
	// }
}
