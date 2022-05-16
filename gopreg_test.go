package goverify

import (
	"fmt"
	"testing"
)

func TestChinaMobile(t *testing.T) {
	for i := 1; i <= 1; i++ {
		for ii := 0; ii <= 9; ii++ {
			for iii := 0; iii <= 9; iii++ {
				one := i
				two := ii
				three := iii
				number := fmt.Sprintf("%d%d%d00138000", one, two, three)
				status, operator := ChinaMobile(number)
				t.Logf("[%s]%s 状态：%v", operator, number, status)
			}
		}
	}
}

func BenchmarkChinaMobile(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func TestChinaMobileNumber(t *testing.T) {
	for i := 1; i <= 1; i++ {
		for ii := 0; ii <= 9; ii++ {
			for iii := 0; iii <= 9; iii++ {
				one := i
				two := ii
				three := iii
				number := fmt.Sprintf("%d%d%d00138000", one, two, three)
				t.Logf("[中国移动]%s 状态：%v", number, ChinaMobileNumber(number))
			}
		}
	}
}

func TestChinaUnicomNumber(t *testing.T) {
	for i := 1; i <= 1; i++ {
		for ii := 0; ii <= 9; ii++ {
			for iii := 0; iii <= 9; iii++ {
				one := i
				two := ii
				three := iii
				number := fmt.Sprintf("%d%d%d00138000", one, two, three)
				t.Logf("[中国联通]%s 状态：%v", number, ChinaUnicomNumber(number))
			}
		}
	}
}

func TestChinaTelecomNumber(t *testing.T) {
	for i := 1; i <= 1; i++ {
		for ii := 0; ii <= 9; ii++ {
			for iii := 0; iii <= 9; iii++ {
				one := i
				two := ii
				three := iii
				number := fmt.Sprintf("%d%d%d00138000", one, two, three)
				t.Logf("[中国电信]%s 状态：%v", number, ChinaTelecomNumber(number))
			}
		}
	}
}

func TestChinaVirtualNumber(t *testing.T) {
	for i := 1; i <= 1; i++ {
		for ii := 0; ii <= 9; ii++ {
			for iii := 0; iii <= 9; iii++ {
				one := i
				two := ii
				three := iii
				number := fmt.Sprintf("%d%d%d00138000", one, two, three)
				t.Logf("[虚拟运营商]%s 状态：%v", number, ChinaVirtualNumber(number))
			}
		}
	}
}
