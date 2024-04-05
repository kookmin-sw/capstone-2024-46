package timeutil

import "time"

var kst, _ = time.LoadLocation("Asia/Seoul")

// ToKST 함수는 주어진 시간을 KST 시간으로 변환합니다.
func ToKST(t time.Time) time.Time {
	return t.In(kst)
}
