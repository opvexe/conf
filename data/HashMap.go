package data

//哈希表
type Table struct {
	array []*TableItem
}

//结构体
type TableItem struct {
	key interface{}
	value interface{}
	table *Table
}

//初始化
func New(size int) *Table {
	return &Table{
		array:make([]*TableItem,size),
	}
}


//赋值
func (t *Table)Set(key int,value interface{})  {
	hash := key%len(t.array)
	if t.array[hash] !=nil {
		if t.array[hash].key!=key {
			//出现碰撞，数据项的table参数指向新的哈希表，但是新哈希表的大小不会变化
			if t.array[hash].table == nil {
				t.array[hash].table = New(len(t.array))
			}
			t.array[hash].table.Set(key,value)
		}else {
			t.array[hash].value = value
		}

	}else {	//当前链表只有一个节点，未赋值
		t.array[hash] = &TableItem{
			key:key,
			value:value,
		}
	}
}

//查询
func (t *Table)Get(key int) interface{} {
	hash := key%len(t.array)
	if t.array[hash]!=nil {
		if t.array[hash].key != key {
			return t.array[hash].table.Get(key)
		} else {
			return t.array[hash].value
		}
	}
	return nil
}