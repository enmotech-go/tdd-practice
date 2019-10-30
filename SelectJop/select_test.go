package SelectJop

import "testing"

func TestRacer(t *testing.T)  {
	slowUrl := "www.baidu.com"
	fastUrl := "www.163.com"
	want := fastUrl
	got := Racer(slowUrl,fastUrl)

	if got !=want{
		t.Errorf("got '%s,want '%s'",got,want)
	}
}
