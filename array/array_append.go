/**
  * Author: JeffreyBool
  * Date: 2019/4/17
  * Time: 16:16
  * Software: GoLand
*/

package array

import (
	"fmt"
)

/**
 * arr 底层扩容知识点
 */
func ArrayAppend() []int {
	arr := make([]int,5)
	fmt.Printf("arr.len: %d; arr.cap: %d \n", len(arr),cap(arr))
	arr = append(arr,10)
	//问现在 arr 结果是什么
	fmt.Printf("arr.len: %d; arr.cap: %d \n", len(arr),cap(arr))
	return arr
}
