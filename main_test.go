package main

import (
	"fmt"
	"testing"
	"time"
)

// toDate yyyy-MM-ddの文字列をtime.Time型に変換するヘルパー関数
func toDate(t *testing.T, dateStr string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", dateStr))
	if err != nil {
		t.Fatalf("toDate() failed (err=%+v)", err)
	}

	return d
}

func TestCreateDateStrs(t *testing.T) {
	var (
		d20210601 = "2021/6/1"
		d20210602 = "2021/6/2"
		d20210603 = "2021/6/3"
	)

	cases := map[string]struct {
		startDate string
		endDate   string
		wants     []*string
		expectErr error
	}{
		"開始日と終了日の差分が0日のケース": {"2021-06-01", "2021-06-01", []*string{&d20210601}, nil},
		"開始日と終了日の差分が2日のケース": {"2021-06-01", "2021-06-03", []*string{&d20210601, &d20210602, &d20210603}, nil},
	}

	for testName, tt := range cases {
		t.Run(testName, func(t * testing.T) {
			t.Parallel()
			gots, err := CreateDateStrs(toDate(t, tt.startDate), toDate(t, tt.endDate))
			if tt.expectErr != nil {
				if err == nil {
					t.Error("want err, but was nil")
				}
			} else {
				if err != nil {
					t.Errorf("not want err (err=%+v)", err)
				}
			}

			if len(gots) != len(tt.wants) {
				t.Errorf("wants length=%+v, but gots length=%+v", len(tt.wants), len(gots))
			}

			for n, got := range gots {
				want := *tt.wants[n]
				if *got != want {
					t.Errorf("want=%+v, but got=%+v", want, *got)
				}
			}
		})
	}
}
