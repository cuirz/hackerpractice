package swordoffer

//剑指 Offer 09. 用两个栈实现队列
//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//提示：
//
//1 <= values <= 10000
//最多会对 appendTail、deleteHead 进行 10000 次调用

type CQueue struct {
	In  []int
	Out []int
}

func Constructor() CQueue {
	return CQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}

}

func (this *CQueue) AppendTail(value int) {
	this.In = append(this.In, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Out) < 1 {
		n := len(this.In)
		if n < 1 {
			return -1
		}
		for i := n - 1; i > -1; i-- {
			this.Out = append(this.Out, this.In[i])
		}
		this.In = make([]int, 0)
	}
	n := len(this.Out)
	res := this.Out[n-1]
	this.Out = this.Out[:n-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
