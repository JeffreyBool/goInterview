/**
  * Author: JeffreyBool
  * Date: 2019/4/17
  * Time: 16:35
  * Software: GoLand
*/

package array

import (
	"testing"
)

func TestArrayEvenNumber(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,12,14,16,19}
	array  := ArrayEvenNumber(arr)
	t.Logf("arr: %+v \n",array)
}
