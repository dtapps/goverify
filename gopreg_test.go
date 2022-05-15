package goverify

import "testing"

func TestMobile(t *testing.T) {
	t.Log(Mobile("13800138000"))
	t.Log(Mobile("33800138000"))
	t.Log(Mobile("1380013800"))
}
