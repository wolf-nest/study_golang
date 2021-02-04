# Map 声明


# Map 元素的访问
1. 在访问的Key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在

# Map 遍历
* map 遍历使用 for range 进行遍历 如下:
```$go
package map_test

import "testing"

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
```