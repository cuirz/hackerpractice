package main

//1603. 设计停车系统
//请你给一个停车场设计一个停车系统。停车场总共有三种不同大小的车位：大，中和小，每种尺寸分别有固定数目的车位。
//
//请你实现ParkingSystem类：
//
//ParkingSystem(int big, int medium, int small)初始化ParkingSystem类，三个参数分别对应每种停车位的数目。
//bool addCar(int carType)检查是否有carType对应的停车位。carType有三种类型：大，中，小，分别用数字1，2和3表示。一辆车只能停在carType对应尺寸的停车位中。如果没有空车位，请返回false，否则将该车停入车位并返回true。
//
//
//示例 1：
//
//输入：
//["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
//[[1, 1, 0], [1], [2], [3], [1]]
//输出：
//[null, true, true, false, false]
//
//解释：
//ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
//parkingSystem.addCar(1); // 返回 true ，因为有 1 个空的大车位
//parkingSystem.addCar(2); // 返回 true ，因为有 1 个空的中车位
//parkingSystem.addCar(3); // 返回 false ，因为没有空的小车位
//parkingSystem.addCar(1); // 返回 false ，因为没有空的大车位，唯一一个大车位已经被占据了
//
//
//提示：
//
//0 <= big, medium, small <= 1000
//carType取值为1，2或3
//最多会调用addCar函数1000次

type ParkingSystem struct {
	Pos [][]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[][]int{nil, {0, big}, {0, medium}, {0, small}}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Pos[carType][0] < this.Pos[carType][1] {
		this.Pos[carType][0]++
		return true
	}
	return false

}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
