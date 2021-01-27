package order

//func bubbleSort(data interface{}, less func(i int, j  int) bool) {
//	dataType := reflect.TypeOf(data).Kind()
//	if dataType == reflect.Slice || dataType == reflect.Array {
//		array := reflect.ValueOf(data)
//		for i := 0; i < array.Len(); i++ {
//			for j := 0; j < array.Len() - i - 1; j++ {
//				if less(j, j+1) {
//					array.Index(j).Set(array.Index(j+1))
//					array.Index(j + 1).Set(array.Index(j))
//					data = array
//				}
//			}
//		}
//	}
//
//}