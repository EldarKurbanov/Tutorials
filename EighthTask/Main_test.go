package EighthTask

import (
	"golang.org/x/tour/tree"
	"reflect"
	"sort"
	"testing"
)

// Упражнение: равнозначные двоичные деревья

func TestWalk(t *testing.T) {
	type args struct {
		t  *tree.Tree
		ch chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{tree.New(1), make(chan int)}},
		{"test2", args{tree.New(2), make(chan int)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "test1":
				treeArray := make([]int, 10)
				go Walk(tt.args.t, tt.args.ch)
				for i := 0; i < 10; i++ {
					num := <- tt.args.ch
					treeArray[i] = num
				}
				sort.Ints(treeArray)
				treeExampleArray := []int{1,2,3,4,5,6,7,8,9,10}
				if !reflect.DeepEqual(treeArray, treeExampleArray) {
					t.Error("Wrong walk when k = 1!")
				}
			case "test2":
				treeArray := make([]int, 10)
				go Walk(tt.args.t, tt.args.ch)
				for i := 0; i < 10; i++ {
					num := <- tt.args.ch
					treeArray[i] = num
				}
				sort.Ints(treeArray)
				treeExampleArray := []int{2,4,6,8,10,12,14,16,18,20}
				if !reflect.DeepEqual(treeArray, treeExampleArray) {
					t.Error("Wrong walk when k = 2!")
				}
			}
		})
	}
}