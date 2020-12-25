package img

import "testing"

func Test_Change(t *testing.T) {
	m, err := Open(`test.jpg`)
	if err != nil {
		t.Fatal("fail to open the file", err)
	}
	m1 := ToGray(m)
	SaveBMP(m1, "test_Gray.bmp")
	m2 := ToBlock(m1)
	SaveBMP(m2, "test_Block.bmp")
	m3 := ToHalftone(m1)
	SaveBMP(m3, "test_Halftone.bmp")

}
