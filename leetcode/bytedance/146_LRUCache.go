package main

//146. LRU缓存机制
//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//
//进阶:
//
//你是否可以在O(1) 时间复杂度内完成这两种操作？
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

type LRUCache struct {
	Array map[int]*Node
	Cap   int
	Size  int
	Head  *Node
	Tail  *Node
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	this := LRUCache{Cap: capacity, Size: 0}
	this.Array = make(map[int]*Node)
	this.Head = &Node{}
	this.Tail = &Node{}
	this.Head.Next = this.Tail
	this.Tail.Pre = this.Head
	return this
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Array[key]; ok {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		node.Pre = this.Head
		node.Next = this.Head.Next
		this.Head.Next.Pre = node
		this.Head.Next = node
		return node.Val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Array[key]; ok {
		node.Val = value
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre

		node.Pre = this.Head
		node.Next = this.Head.Next
		this.Head.Next.Pre = node
		this.Head.Next = node
	} else {
		if this.Size == this.Cap {
			tail := this.Tail.Pre
			tail.Pre.Next = this.Tail
			this.Tail.Pre = tail.Pre
			delete(this.Array, tail.Key)
		} else {
			this.Size++
		}

		node = &Node{Key: key, Val: value}
		this.Array[key] = node

		node.Pre = this.Head
		node.Next = this.Head.Next
		this.Head.Next.Pre = node
		this.Head.Next = node
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	obj.Get(1)
	obj.Get(2)

	//["LRUCache","put","put","put","put","get","get"]
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	//[null,null,null,1,null,2,null,-1,-1,4]
	//[null,null,null,1,null,-1,null,-1,3,4]

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	//["LRUCache","put","get","put","get","get"]
	//[[1],[2,1],[2],[3,2],[2],[3]]

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
}
