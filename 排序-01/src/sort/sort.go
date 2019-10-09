package sort

//冒泡排序
func BobbleSort(arr []int) {
	falg := false
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//相邻元素比较大小
			if arr[j] > arr[j+1] {
				//交换元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
				falg = true
			}
		}
		if !falg {
			return
		} else {
			falg = true
		}
	}
}

//选择排序
func SelectSort(arr []int) {
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//定义变量保存最大值
		index := 0
		//内层控制列
		for j := 1; j < len(arr)-i; j++ { //0与自己比较无意义
			//元素比较大小
			if arr[j] > arr[index] {
				index = j
			}
		}
		if index != 0 {
			arr[index], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[index]
		}
	}
}

//插入排序
func InsertSort(arr []int) {
	//假设 第一个元素arr[0] 作为有序组，循环依次取出无序组元素
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] { //将第一个无序组元素与最后一个有序组元素比较
			for j := i; j > 0; j-- { //有序组  循环比较有序组中的每一个元素，确定最终位置
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				} else {
					break
				}
			}
		}
	}
}

//希尔排序
func ShellSort(arr []int) {
	//利用循环去修改增量值 len()arr/2  .... len()arr/2/2
	for inc := len(arr) / 2; inc > 0; inc /= 2 {
		//利用增量，做循环的起始值，依次后移取出元素
		for i := inc; i < len(arr); i++ {
			//定义变量
			temp := arr[i]
			//利用增量数据去找到对应元素
			for j := i - inc; j >= 0; j -= inc {
				if temp < arr[j] {
					arr[j], arr[j+inc] = arr[j+inc], arr[j]
				} else {
					break
				}
			}
		}
	}
}

//快速排序
func QicklySort(arr []int, start int, end int,n int) {
	//起始下标小于结束下标
	if start < end {
		i, j, base := start, end, arr[start]
		//i,j对向移动，直至重合
		for i < j {
			//j 先判断后移动,大于基准值前移
			for i < j && arr[j] >= base {
				j--
			}
			//小于基准值，与i对应的元素交换
			arr[i], arr[j] = arr[j], arr[i]

			for i < j && arr[i] <= base {
				i++
			}
			//大于基准值，与j对应的元素交换
			arr[i], arr[j] = arr[j], arr[i]
		}
		//递归去调用本函数
		QicklySort(arr, start, i-1,1)
		QicklySort(arr, i+1, end,2)
	}
}

/*
 调用
 QicklySort(arr,0,len(arr)-1)
*/

//堆排
func HeapSort(arr []int) {
	//将二叉树调整成最大堆存储，保证各个子树,(3个树种的最大值在根位置)
	length := len(arr)                   //9
	for i := length/2 - 1; i >= 0; i-- { //3,2,1,0  是各个"根"
		CreateMaxHeap(arr, i, length-1)
	}
	//上述for 循环，不从书上摘下节点

	//此循环开始，根节点保存整个数组的最大值
	for i:=length-1;i>0;i--{
		//如果只剩根节点和左子节点，排序结束
		if i==1&&arr[0]<=arr[i] {
			break
		}
		//将根节点和最后一个叶子节点交换
		arr[0],arr[i] = arr[i],arr[0]
		//将交换后的最后一个节点摘下，确认是最大的，无需在比较
		CreateMaxHeap(arr,0,i-1)
	}
}

//获取堆中最大值 （3，8） （2，8）...
func CreateMaxHeap(arr []int, startNode int, maxNode int) {
	//存放较大值的下标
	var max int
	//定义坐姿节点下标和右子节点下标
	lChild := startNode*2 + 1		//7 ,5
	rChild := lChild + 1		//8,6

	//左子节点下标查出最大下标，跳出递归
	if lChild >= maxNode {
		return
	}
	//左右比较，找出最大值
	if rChild <= maxNode && arr[rChild] > arr[lChild] {
		max = rChild
	} else {
		max = lChild
	}

	//与根节点比较
	if arr[max] <= arr[startNode] {
		return
	}
	//比根节点大，交换数据
	arr[startNode], arr[max] = arr[max], arr[startNode]
	//递归进行下次比较
	CreateMaxHeap(arr, max, maxNode)
}

//二分法查找				//前提是有序的数
func BinarySearch(arr []int, num int) int {
	//定义起始下标
	start := 0
	//定义结尾下标
	end := len(arr) - 1
	//定义中间位置
	mid := (start + end) / 2
	//循环查找
	for i := 0; i < len(arr); i++ {
		if num == arr[mid] {
			return mid
		} else if num > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	return -1
}
