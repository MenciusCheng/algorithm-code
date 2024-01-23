package mergeflight_example

import "testing"

func TestChannelParamsForTrap(t *testing.T) {
	ChannelParamsForTrap()
}

func TestChannelParamsForFix(t *testing.T) {
	n := 1000
	res := ChannelParamsForFix(n)
	if n != len(res) {
		t.Errorf("want len %d, got len %d", n, len(res))
		return
	}
	t.Logf("res = %+v", res)
}
