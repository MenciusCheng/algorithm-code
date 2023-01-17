package w20230107

import "testing"

func TestDataStream_Consec(t *testing.T) {
	dataStream := Constructor(4, 3) // value = 4, k = 3
	if got := dataStream.Consec(4); got != false {
		t.Errorf("Consec() = %v, want %v", got, false)
	} // 数据流中只有 1 个整数，所以返回 False 。
	if got := dataStream.Consec(4); got != false {
		t.Errorf("Consec() = %v, want %v", got, false)
	} // 数据流中只有 2 个整数，由于 2 小于 k ，返回 False 。
	if got := dataStream.Consec(4); got != true {
		t.Errorf("Consec() = %v, want %v", got, false)
	} // 数据流最后 3 个整数都等于 value， 所以返回 True 。
	if got := dataStream.Consec(3); got != false {
		t.Errorf("Consec() = %v, want %v", got, false)
	} // 最后 k 个整数分别是 [4,4,3] 。
}
