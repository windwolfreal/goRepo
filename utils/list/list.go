package utils

import (
	orgList  "container/list"
)

type List orgList.List

func (source *List) ToSlice() []interface{} {
	oli := (*orgList.List)(source)
	
	length :=oli.Len()
	slice := make([]interface{}, length)
	for item, i := oli.Front(), 0; item != nil;  {
		slice[i] = item
		item = item.Next()
		i++
	}
	
	return slice
}