package medianFinder

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSec(t *testing.T) {
	a := []int{78, 14, 50, 20, 13, 9, 25, 8, 13, 37, 29, 33, 55, 52, 6, 17, 65, 23, 74, 43, 5, 29, 29, 72, 7, 13, 56, 21, 31, 66, 69, 69, 74, 12, 77, 23, 10, 6, 27, 63, 77, 21, 40, 10, 19, 59, 35, 40}
	obj := Constructor()
	for _, value := range a {
		obj.AddNum(value)
	}
	t.Log(obj.FindMedian())

}

func TestFirst(t *testing.T) {
	funcNameLists := []string{"addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	argLists := [][]int{{78}, {}, {14}, {}, {50}, {}, {20}, {}, {13}, {}, {9}, {}, {25}, {}, {8}, {}, {13}, {}, {37}, {}, {29}, {}, {33}, {}, {55}, {}, {52}, {}, {6}, {}, {17}, {}, {65}, {}, {23}, {}, {74}, {}, {43}, {}, {5}, {}, {29}, {}, {29}, {}, {72}, {}, {7}, {}, {13}, {}, {56}, {}, {21}, {}, {31}, {}, {66}, {}, {69}, {}, {69}, {}, {74}, {}, {12}, {}, {77}, {}, {23}, {}, {10}, {}, {6}, {}, {27}, {}, {63}, {}, {77}, {}, {21}, {}, {40}, {}, {10}, {}, {19}, {}, {59}, {}, {35}, {}, {40}, {}, {44}, {}, {4}, {}, {15}, {}, {29}, {}, {63}, {}, {27}, {}, {46}, {}, {56}, {}, {0}, {}, {60}, {}, {72}, {}, {35}, {}, {54}, {}, {50}, {}, {14}, {}, {29}, {}, {62}, {}, {24}, {}, {18}, {}, {79}, {}, {16}, {}, {19}, {}, {8}, {}, {77}, {}, {10}, {}, {21}, {}, {66}, {}, {42}, {}, {76}, {}, {14}, {}, {58}, {}, {20}, {}, {0}, {}}

	obj := Constructor()
	var value []reflect.Value
	var args interface{}
	for index, funcName := range funcNameLists {
		funcName = strFirstToUpper(funcName)
		if len(argLists[index]) > 0 {
			args = argLists[index][0]
		} else {
			args = -1
		}
		if args == -1 {
			value = InvokeObjectMethod(&obj, funcName)
		} else {
			value = InvokeObjectMethod(&obj, funcName, args)
		}
		if len(value) > 0 {
			t.Logf("obj.%s()=%f", funcName, value[0].Float())
		} else {
			t.Logf("obj.%s(%d)", funcName, args)
		}
	}

	//funcName := "addNum"
	//funcName = strFirstToUpper(funcName)
	//args := []int{3}
	//obj := Constructor()
	//InvokeObjectMethod(&obj,funcName,args)
	//args = []int{}
	//t.Log(InvokeObjectMethod(&obj,"FindMedian",args)[0].Float())
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) []reflect.Value {
	//fmt.Println(reflect.ValueOf(object))
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}
