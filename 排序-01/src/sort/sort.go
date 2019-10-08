package sort

//冒泡排序
func BobbleSort(arr []int)  {
	falg := false
	//外层控制行
	for i := 0; i<len(arr)-1;i++  {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//相邻元素比较大小
			if arr[j]>arr[j+1] {
				//交换元素
				arr[j],arr[j+1] = arr[j+1],arr[j]
				falg = true
			}
		}
		if !falg {
			return
		}else {
			falg = true
		}
	}
}

//选择排序
func SelectSort(arr []int)  {
	//外层控制行
	for i := 0; i<len(arr)-1;i++  {
		//定义变量保存最大值
		index := 0
		//内层控制列
		for j := 1; j < len(arr)-i; j++ {			//0与自己比较无意义
			//元素比较大小
			if arr[j]>arr[index] {
				index = j
			}
		}
		if index!=0 {
			arr[index],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[index]
		}
	}
}

//插入排序
func InsertSort(arr []int)  {
	//假设 第一个元素arr[0] 作为有序组，循环依次取出无序组元素
	for i:=1;i<len(arr);i++{
		if arr[i]<arr[i-1] {  //将第一个无序组元素与最后一个有序组元素比较
			for j:=i;j>0;j--{ //有序组  循环比较有序组中的每一个元素，确定最终位置
				if arr[j] <arr[j-1]{
					arr[j],arr[j-1] = arr[j-1],arr[j]
				}else {
					break
				}
			}
		}
	}
}

//计数统计排序
func CountingSort()  {

}