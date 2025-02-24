package util_test

import (
	"log"
	"testing"

	u "github.com/bruno5200/xyz/util"
)

func TestLiteral(t *testing.T) {

	var number int = 1997

	literal, err := u.Literal(number)

	if err != nil {
		log.Println(err)
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Number: %d, Lietral: %s", number, literal)
}

func TestLiteralString(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		number int
		want   string
		test   string
		pass   bool
	}{
		{
			1997,
			"mil novecientos noventa y siete",
			"valid number",
			true,
		},
		{
			-1997,
			"menos mil novecientos noventa y siete",
			"valid number",
			true,
		},
		{
			0,
			"cero",
			"valid number",
			true,
		},
		{
			1000000,
			"un millón",
			"valid number",
			true,
		},
		{
			1000001,
			"un millón uno",
			"valid number",
			true,
		},
		{
			1000000000,
			"mil millones",
			"valid number",
			true,
		},
		{
			1000000001,
			"mil millones uno",
			"valid number",
			true,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			literal, err := u.Literal(test.number)
			if (err != nil) == test.pass {
				t.Errorf("unexpected error: %v", err)
			}
			if literal != test.want {
				t.Errorf("got %q, want %q", literal, test.want)
			}
		})
	}
}
