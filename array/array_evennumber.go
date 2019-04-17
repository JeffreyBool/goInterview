/**
  * Author: JeffreyBool
  * Date: 2019/4/17
  * Time: 16:20
  * Software: GoLand
*/

package array


/**
 * 考察点Slices的变量储存方式 (切片是引用类型)
 * 所以每次对 array.append 做的修改,本身会对 array 指针指向的变量地址的值做修改.
 */
func ArrayEvenNumber(array []int) []int {
	for index,arr := range array{
		if arr % 2 == 0 {
			array = append(array[:index],array[index +1:]...)
		}
	}
	return array
}