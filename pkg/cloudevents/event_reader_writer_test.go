package cloudevents_test

import (
	ce "github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
	"time"
)

type ReadWriteTest struct {
	event   ce.Event
	set     string
	want    interface{}
	wantErr string
}

func TestEventRW_SpecVersion(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"empty v01": {
			event: ce.New(""),
			set:   "0.1",
			want:  "0.1",
		},
		"empty v02": {
			event: ce.New(""),
			set:   "0.2",
			want:  "0.2",
		},
		"empty v03": {
			event: ce.New(""),
			set:   "0.3",
			want:  "0.3",
		},
		"v01": {
			event: ce.New("0.1"),
			set:   "0.1",
			want:  "0.1",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "0.2",
			want:  "0.2",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "0.3",
			want:  "0.3",
		},
		"invalid v01": {
			event:   ce.New("0.1"),
			set:     "1.1",
			want:    "0.1",
			wantErr: "invalid version",
		},
		"invalid v02": {
			event:   ce.New("0.2"),
			set:     "1.2",
			want:    "0.2",
			wantErr: "invalid version",
		},
		"invalid v03": {
			event:   ce.New("0.3"),
			set:     "1.3",
			want:    "0.3",
			wantErr: "invalid version",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetSpecVersion(tc.set)
			got := tc.event.SpecVersion()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_Type(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "type.0.1",
			want:  "type.0.1",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "type.0.2",
			want:  "type.0.2",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "type.0.3",
			want:  "type.0.3",
		},
		"spaced v01": {
			event: ce.New("0.1"),
			set:   "  type.0.1  ",
			want:  "type.0.1",
		},
		"spaced v02": {
			event: ce.New("0.2"),
			set:   "  type.0.2  ",
			want:  "type.0.2",
		},
		"spaced v03": {
			event: ce.New("0.3"),
			set:   "   type.0.3   ",
			want:  "type.0.3",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetType(tc.set)
			got := tc.event.Type()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_ID(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "id.0.1",
			want:  "id.0.1",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "id.0.2",
			want:  "id.0.2",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "id.0.3",
			want:  "id.0.3",
		},
		"spaced v01": {
			event: ce.New("0.1"),
			set:   "  id.0.1  ",
			want:  "id.0.1",
		},
		"spaced v02": {
			event: ce.New("0.2"),
			set:   "  id.0.2  ",
			want:  "id.0.2",
		},
		"spaced v03": {
			event: ce.New("0.3"),
			set:   "   id.0.3   ",
			want:  "id.0.3",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetID(tc.set)
			got := tc.event.ID()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_Source(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"invalid v01": {
			event:   ce.New("0.1"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
		"invalid v02": {
			event:   ce.New("0.2"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
		"invalid v03": {
			event:   ce.New("0.3"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetSource(tc.set)
			got := tc.event.Source()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_Subject(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "subject.0.1",
			want:  "subject.0.1",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "subject.0.2",
			want:  "subject.0.2",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "subject.0.3",
			want:  "subject.0.3",
		},
		"spaced v01": {
			event: ce.New("0.1"),
			set:   "  subject.0.1  ",
			want:  "subject.0.1",
		},
		"spaced v02": {
			event: ce.New("0.2"),
			set:   "  subject.0.2  ",
			want:  "subject.0.2",
		},
		"spaced v03": {
			event: ce.New("0.3"),
			set:   "   subject.0.3   ",
			want:  "subject.0.3",
		},
		"nilled v01": {
			event: func() ce.Event {
				e := ce.New("0.1")
				_ = e.SetSource("should nil")
				return e
			}(),
			want: "",
		},
		"nilled v02": {
			event: func() ce.Event {
				e := ce.New("0.2")
				_ = e.SetSource("should nil")
				return e
			}(),
			want: "",
		},
		"nilled v03": {
			event: func() ce.Event {
				e := ce.New("0.3")
				_ = e.SetSource("should nil")
				return e
			}(),
			want: "",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetSubject(tc.set)
			got := tc.event.Subject()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_Time(t *testing.T) {
	now := time.Now()

	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "now", // hack
			want:  now,
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "now", // hack
			want:  now,
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "now", // hack
			want:  now,
		},
		"nilled v01": {
			event: func() ce.Event {
				e := ce.New("0.1")
				_ = e.SetTime(now)
				return e
			}(),
			want: time.Time{},
		},
		"nilled v02": {
			event: func() ce.Event {
				e := ce.New("0.2")
				_ = e.SetTime(now)
				return e
			}(),
			want: time.Time{},
		},
		"nilled v03": {
			event: func() ce.Event {
				e := ce.New("0.3")
				_ = e.SetTime(now)
				return e
			}(),
			want: time.Time{},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			var err error
			if tc.set == "now" {
				err = tc.event.SetTime(now) // pull now from outer test.
			} else {
				err = tc.event.SetTime(time.Time{}) // pull now from outer test.
			}
			got := tc.event.Time()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_SchemaURL(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "http://example/",
			want:  "http://example/",
		},
		"invalid v01": {
			event:   ce.New("0.1"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
		"invalid v02": {
			event:   ce.New("0.2"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
		"invalid v03": {
			event:   ce.New("0.3"),
			set:     "%",
			want:    "",
			wantErr: "invalid URL escape",
		},
		"nilled v01": {
			event: func() ce.Event {
				e := ce.New("0.1")
				_ = e.SetSchemaURL("should nil")
				return e
			}(),
			want: "",
		},
		"nilled v02": {
			event: func() ce.Event {
				e := ce.New("0.2")
				_ = e.SetSchemaURL("should nil")
				return e
			}(),
			want: "",
		},
		"nilled v03": {
			event: func() ce.Event {
				e := ce.New("0.3")
				_ = e.SetSchemaURL("should nil")
				return e
			}(),
			want: "",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetSchemaURL(tc.set)
			got := tc.event.SchemaURL()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_DataContentType(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "application/json",
			want:  "application/json",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "application/json",
			want:  "application/json",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "application/json",
			want:  "application/json",
		},
		"spaced v01": {
			event: ce.New("0.1"),
			set:   "  application/json  ",
			want:  "application/json",
		},
		"spaced v02": {
			event: ce.New("0.2"),
			set:   "  application/json  ",
			want:  "application/json",
		},
		"spaced v03": {
			event: ce.New("0.3"),
			set:   "   application/json   ",
			want:  "application/json",
		},
		"nilled v01": {
			event: func() ce.Event {
				e := ce.New("0.1")
				_ = e.SetDataContentType("application/json")
				return e
			}(),
			want: "",
		},
		"nilled v02": {
			event: func() ce.Event {
				e := ce.New("0.2")
				_ = e.SetDataContentType("application/json")
				return e
			}(),
			want: "",
		},
		"nilled v03": {
			event: func() ce.Event {
				e := ce.New("0.3")
				_ = e.SetDataContentType("application/json")
				return e
			}(),
			want: "",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetDataContentType(tc.set)
			got := tc.event.DataContentType()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func TestEventRW_DataContentEncoding(t *testing.T) {
	testCases := map[string]ReadWriteTest{
		"v01": {
			event: ce.New("0.1"),
			set:   "base64",
			want:  "base64",
		},
		"v02": {
			event: ce.New("0.2"),
			set:   "base64",
			want:  "base64",
		},
		"v03": {
			event: ce.New("0.3"),
			set:   "base64",
			want:  "base64",
		},
		"spaced v01": {
			event: ce.New("0.1"),
			set:   "  base64  ",
			want:  "base64",
		},
		"spaced v02": {
			event: ce.New("0.2"),
			set:   "  base64  ",
			want:  "base64",
		},
		"spaced v03": {
			event: ce.New("0.3"),
			set:   "   base64   ",
			want:  "base64",
		},
		"cased v01": {
			event: ce.New("0.1"),
			set:   "  BaSe64  ",
			want:  "base64",
		},
		"cased v02": {
			event: ce.New("0.2"),
			set:   "  BaSe64  ",
			want:  "base64",
		},
		"cased v03": {
			event: ce.New("0.3"),
			set:   "   BaSe64   ",
			want:  "base64",
		},
		"nilled v01": {
			event: func() ce.Event {
				e := ce.New("0.1")
				_ = e.SetDataContentEncoding("base64")
				return e
			}(),
			want: "",
		},
		"nilled v02": {
			event: func() ce.Event {
				e := ce.New("0.2")
				_ = e.SetDataContentEncoding("base64")
				return e
			}(),
			want: "",
		},
		"nilled v03": {
			event: func() ce.Event {
				e := ce.New("0.3")
				_ = e.SetDataContentEncoding("base64")
				return e
			}(),
			want: "",
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			err := tc.event.SetDataContentEncoding(tc.set)
			got := tc.event.DataContentEncoding()

			validateReaderWriter(t, tc, got, err)
		})
	}
}

func validateReaderWriter(t *testing.T, tc ReadWriteTest, got interface{}, err error) {
	var gotErr string
	if err != nil {
		gotErr = err.Error()
		if tc.wantErr == "" {
			t.Errorf("unexpected no error, got %q", gotErr)
		}
	}
	if tc.wantErr != "" {
		if !strings.Contains(gotErr, tc.wantErr) {
			t.Errorf("unexpected error, expected to contain %q, got: %q ", tc.wantErr, gotErr)
		}
	}
	if diff := cmp.Diff(tc.want, got); diff != "" {
		t.Errorf("unexpected (-want, +got) = %v", diff)
	}
}