package chapter

// 此题选择用java写，因为大整数精度没法用做map的key

// import "math/big"

// func maxPoints(points [][]int) int {
// 	result := 0
// 	// 针对每一个点
// 	// case1: same point; case2: same x; case 3: k
// 	for i := 0; i < len(points); i++ {
// 		samePoint := 0
// 		slopeMap := make(map[big.Rat]int)
// 		for j := i + 1; j < len(points); j++ {
// 			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
// 				samePoint++
// 				continue
// 			}
// 		}
// 	}
// 	return result
// }
