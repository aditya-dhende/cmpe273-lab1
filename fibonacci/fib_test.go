package fibonacci

import "testing"

func TestFibSeries(t *testing.T) {

	var num int
	num = fibSeries(10)
	if num != 55 {
		t.Error("Expected 55, got", num)
	}
  
    


}

func TestFibSeries2(t *testing.T) {

	var num int
	num = fibSeries(12)
	if num != 144 {
		t.Error("Expected 144, got", num)
	}
  

}

