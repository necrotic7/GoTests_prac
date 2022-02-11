package main

import (
	
	"reflect"
	"testing"
)

func TestSummer(t *testing.T){

	t.Run("collection of any size", func(t *testing.T){
		numbers := []int{
			1, 2, 3,
		}
		got := Summer(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
	})
	
}

func TestSumAll(t *testing.T){
	got := SumAll([]int{1, 2}, []int{0, 9})
	
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCollectAllTail(t *testing.T){

	t.Run("collect tails", func(t *testing.T){
		got := CollectAllTail([]int{6}, []int{0, 9, 8})
		want := []int{6, 8}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})
	
	t.Run("safely sum empty slice", func(t *testing.T){
		
		got := CollectAllTail([]int{}, []int{7, 8})
		want := []int{8}
		
		assertCorrectMessage(t, got, want)
		
	})

}