package main

import(
	"testing"
)

func CheckPerimeter(t *testing.T, shape Shape, want float64){
	t.Helper()
	got := shape.Perimeter()
	assertCorrectMessage(t, got, want)
}

func CheckArea(t *testing.T, shape Shape, want float64){
	t.Helper()
	got := shape.Area()
	assertCorrectMessage(t, got, want)
}

func TestPerimeter(t *testing.T){

	t.Run("caculating Rectangle", func(t *testing.T){
		r := Rectangle{10.0, 10.0}
		want := 40.0
		CheckPerimeter(t, r, want)
	})
	
	t.Run("caculating Circle", func(t *testing.T){
		c := Circle{2.0}
		want := 12.56
		CheckPerimeter(t, c, want)
	})

	t.Run("caculating Triangle", func(t *testing.T){
		tri := Triangle{5.0, 5.0}
		want := 15.0
		CheckPerimeter(t, tri, want)
	})
}

func TestArea(t *testing.T){
	//TableDrivenTests
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{2.0}, 12.56},
		{Triangle{6.0, 6.0}, 18.0},
	}

	for _, test := range(areaTests){
		CheckArea(t, test.shape, test.want)
	}
}