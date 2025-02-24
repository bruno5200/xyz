package util_test

import (
	"testing"
	"time"

	u "github.com/bruno5200/xyz/util"
	"github.com/google/uuid"
)

func TestStringToInt(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want int
		test string
	}{
		{
			"",
			0,
			"empty string",
		},
		{
			"1",
			1,
			"valid string to int",
		},
		{
			"2b",
			0,
			"invalid number",
		},
		{
			"2b%",
			0,
			"invalid string with special characters",
		},
		{
			"A2",
			0,
			"invalid string with letters",
		},
		{
			"1000",
			1000,
			"valid",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		got, err := u.StringToInt(test.s)
		if err != nil {
			t.Logf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
		}
		if got != test.want {
			t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
		}
	}
}

func TestStringtoInt64(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want int64
		test string
	}{
		{
			"",
			0,
			"empty string",
		},
		{
			"1",
			1,
			"valid string to int",
		},
		{
			"2b",
			0,
			"invalid number",
		},
		{
			"2b%",
			0,
			"invalid string with special characters",
		},
		{
			".",
			0,
			"invalid string with point",
		},
		{
			"2.0",
			20,
			"invalid string number with point",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			got := u.StringToInt64(test.s)
			if got != test.want {
				t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
			}
			t.Logf("Result %d", got)
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want float64
		test string
	}{
		{
			"",
			0,
			"empty string",
		},
		{
			"1",
			1,
			"valid string to int",
		},
		{
			"2,18",
			2.18,
			"valid string to float",
		},
		{
			"2.1°",
			0,
			"invalid decimal with letters",
		},
		{
			"2b",
			0,
			"invalid number",
		},
		{
			"2b%",
			0,
			"invalid string with special characters",
		},
		{
			".",
			0,
			"invalid string with point",
		},
		{
			"2.0",
			2.0,
			"invalid string number with point",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			got, _ := u.StringToFloat64(test.s)
			if got != test.want {
				t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
			}
			t.Logf("Result %f", got)
		})
	}
}

func TestPointerToString(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    *string
		want string
		test string
	}{
		{
			nil,
			"",
			"empty string",
		},
		{
			stringToPointer(""),
			"",
			"empty string",
		},
		{
			stringToPointer("test"),
			"test",
			"valid string",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			got := u.PointerToString(test.s)
			if got != test.want {
				t.Errorf("StringToInt(%v) = %v, want %v", test.s, got, test.want)
			}
			t.Logf("Result %s", got)
		})
	}
}

func TestUnixTimestampStringToDate(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want string
		test string
	}{
		{
			"",
			"",
			"empty string",
		},
		{
			"632721600000",
			"19/01/1990",
			"valid string to date",
		},
		{
			"2b",
			"",
			"invalid number",
		},
		{
			"2b%",
			"",
			"invalid string with special characters",
		},
		{
			".",
			"",
			"invalid string with point",
		},
		{
			"2.0",
			"",
			"invalid string number with point",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			got := u.UnixTimestampStringToDate(test.s)
			if got != test.want {
				t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
			}
			t.Logf("Result %s", got)
		})
	}
}

func TestFormatDateWithYear(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		date string
		want string
	}{
		{
			date: "21/12/2018",
			want: "21 de Diciembre, 2018",
		},
		{
			date: "14/11/2023",
			want: "14 de Noviembre, 2023",
		},
		{
			date: "25/10/2023",
			want: "25 de Octubre, 2023",
		},
		{
			date: "17/09/2023",
			want: "17 de Septiembre, 2023",
		},
		{
			date: "10/08/1997",
			want: "10 de Agosto, 1997",
		},
		{
			date: "01/07/2004",
			want: "01 de Julio, 2004",
		},
		{
			date: "02/06/2005",
			want: "02 de Junio, 2005",
		},
		{
			date: "03/05/2006",
			want: "03 de Mayo, 2006",
		},
		{
			date: "04/04/2007",
			want: "04 de Abril, 2007",
		},
		{
			date: "05/03/2008",
			want: "05 de Marzo, 2008",
		},
		{
			date: "06/02/2009",
			want: "06 de Febrero, 2009",
		},
		{
			date: "07/01/2010",
			want: "07 de Enero, 2010",
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.date, func(t *testing.T) {
			t.Parallel()

			date, _ := time.Parse("02/01/2006", test.date)
			got := u.FormatDateWithYear(date)
			t.Logf("Testing %s", test.date)
			if got != test.want {
				t.Errorf("FormatDate(%s) = %s, want %s", test.date, got, test.want)
			}
		})
	}
}

func TestFormatDateWithoutYear(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		date string
		want string
	}{
		{
			date: "21/12/2018",
			want: "21 de Diciembre",
		},
		{
			date: "14/11/2023",
			want: "14 de Noviembre",
		},
		{
			date: "25/10/2023",
			want: "25 de Octubre",
		},
		{
			date: "17/09/2023",
			want: "17 de Septiembre",
		},
		{
			date: "10/08/1997",
			want: "10 de Agosto",
		},
		{
			date: "01/07/2004",
			want: "01 de Julio",
		},
		{
			date: "02/06/2005",
			want: "02 de Junio",
		},
		{
			date: "03/05/2006",
			want: "03 de Mayo",
		},
		{
			date: "04/04/2007",
			want: "04 de Abril",
		},
		{
			date: "05/03/2008",
			want: "05 de Marzo",
		},
		{
			date: "06/02/2009",
			want: "06 de Febrero",
		},
		{
			date: "07/01/2010",
			want: "07 de Enero",
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.date, func(t *testing.T) {
			t.Parallel()

			date, _ := time.Parse("02/01/2006", test.date)
			got := u.FormatDateWithoutYear(date)
			t.Logf("Testing %s", test.date)
			if got != test.want {
				t.Errorf("FormatDate(%s) = %s, want %s", test.date, got, test.want)
			}
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	date := time.Now()
	t.Log(u.FormatDateTime(date))
	t.Log(date.Format("2006-01-02 15:04:05+00"))
}

func TestFirstIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "111a9044",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "5192a55d",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "81617e7a",
		},
	}
	for _, test := range tests {
		got := u.FisrtIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestSecondIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "33c9",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "1db0",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "ce55",
		},
	}
	for _, test := range tests {
		got := u.SecondIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestThirdIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "4306",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "4042",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "42ae",
		},
	}
	for _, test := range tests {
		got := u.ThirdIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestFourthIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "9feb",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "bd28",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "b440",
		},
	}
	for _, test := range tests {
		got := u.FourthIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestFifthIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "f7501cabca29",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "ab9b6b2e7744",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "ea6916ae943a",
		},
	}
	for _, test := range tests {
		got := u.FifthIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func stringToPointer(s string) *string {
	return &s
}
