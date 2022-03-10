package main

//460. LFU缓存
//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。它应该支持以下操作：get和put。
//
//get(key)- 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
//put(key, value)- 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除最久未使用的键。
//「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
//
//
//
//进阶：
//你是否可以在O(1)时间复杂度内执行两项操作？
//
//
//
//示例：
//
//LFUCache cache = new LFUCache( 2 /* capacity (缓存容量) */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回 1
//cache.put(3, 3);    // 去除 key 2
//cache.get(2);       // 返回 -1 (未找到key 2)
//cache.get(3);       // 返回 3
//cache.put(4, 4);    // 去除 key 1
//cache.get(1);       // 返回 -1 (未找到 key 1)
//cache.get(3);       // 返回 3
//cache.get(4);       // 返回 4

type LFUCache struct {
	// 频率哈希表, key 频率 value 双向链表
	FreqTable map[int]*LinkedList
	// 哈希表
	ArrayTable map[int]*Node
	Size       int
	Cap        int
	MinFreq    int
}

type LinkedList struct {
	Key  int
	Head *Node
	Tail *Node
}

type Node struct {
	Key   int
	Value int
	Count int
	Pre   *Node
	Next  *Node
}

func Constructor(capacity int) LFUCache {
	this := LFUCache{
		FreqTable:  make(map[int]*LinkedList),
		ArrayTable: make(map[int]*Node),
		Cap:        capacity,
	}
	return this

}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.ArrayTable[key]; ok {
		this.removeFromLinked(node)
		node.Count++
		this.addToLinked(node)
		return node.Value
	}
	return -1

}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap == 0 {
		return
	}
	if node, ok := this.ArrayTable[key]; ok {
		node.Value = value
		this.removeFromLinked(node)
		node.Count++
		this.addToLinked(node)
	} else {
		if this.Size == this.Cap {
			// 删除最小调用节点
			this.removeMinFreq()
		} else {
			this.Size++
		}
		this.MinFreq = 1
		count := 1
		node := &Node{Key: key, Count: count, Value: value}
		this.addToLinked(node)
		this.ArrayTable[key] = node
	}
}

func removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LinkedList) addToHead(node *Node) {
	node.Pre = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LFUCache) removeMinFreq() {
	if linked, has := this.FreqTable[this.MinFreq]; has {
		node := linked.Tail.Pre
		node.Pre.Next = node.Next
		linked.Tail.Pre = node.Pre
		delete(this.ArrayTable, node.Key)
	}
}

func (this *LFUCache) removeFromLinked(node *Node) {
	removeNode(node)
	if linked, has := this.FreqTable[node.Count]; has {
		if linked.Head.Next == linked.Tail && node.Count == this.MinFreq {
			this.MinFreq++
		}
	}
}
func (this *LFUCache) addToLinked(node *Node) {
	if linked, has := this.FreqTable[node.Count]; has {
		linked.addToHead(node)
	} else {
		linked = &LinkedList{Key: node.Count, Head: &Node{}, Tail: &Node{}}
		linked.Head.Next = linked.Tail
		linked.Tail.Pre = linked.Head
		linked.addToHead(node)
		this.FreqTable[node.Count] = linked
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(0)
	obj.Put(0, 0)
	obj.Get(0)

	//["LFUCache","put","get"]
	//[[0],[0,0],[0]]
	//["LFUCache","put","put","get","put","get","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]

	//[null,null,null,1,null,2,3,null,1,3,4]
	//[null,null,null,1,null,-1,3,null,-1,3,4]
}
