// 测试用函数
package Helper

import (
	"math/rand"
	"time"
)

// 生成元素数量为n,元素在[rangeL,rangeR]区间的 整数数组
func GenerateRandomArray(n int, rangeL int, rangeR int) map[int]int {
	if rangeL > rangeR {
		return nil
	}

	arr := map[int]int{} // rand

	rand2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		// 指定范围随机数生成的标准写法
		arr[i] = rand2.Int()%(rangeR-rangeL+1) + rangeL
	}

	return arr
}

// 生成近似顺序的整数数组
// n: 数组大小; swapTimes: 数组打乱次数;
func GenerateNearlyArray(n int, swapTimes int) map[int]int {
	arr := make(map[int]int)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	rand2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < swapTimes; i++ {
		// 取两个 [0,n) 区间的数
		posX := rand2.Int() % n
		posy := rand2.Int() % n
		// swap 随机交换两个数据
		arr[posX], arr[posy] = arr[posy], arr[posX]
	}

	return arr
}

func MaxInt(n int, m int) int {
	if n > m {
		return n
	}
	return m
}

// 去掉非数字与字母的部分，并将大写字母转换为小写字母
// a-z: 97-122
// A-Z: 65-90
// 0-9: 48-57
func SimplifyStr(s string) string {
	btS := []byte(s)

	var nS []byte
	for i := 0; i < len(btS); i++ {
		if (btS[i] >= 48 && btS[i] <= 57) || (btS[i] >= 97 && btS[i] <= 122) {
			nS = append(nS, btS[i])
		} else if btS[i] >= 65 && btS[i] <= 90 {
			nS = append(nS, btS[i]+32)
		}
	}

	return string(nS)
}
