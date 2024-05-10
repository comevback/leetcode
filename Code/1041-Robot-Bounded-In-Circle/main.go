package main

func main() {

}

// ==============================================================================================================================
// 1. 创建机器人结构以及方法 Create Robot struct and methods
func isRobotBounded(instructions string) bool {
	// 创建机器人对象 Create a robot object
	robot := NewRobot()
	// 首先进行一次循环，如果机器人回到原点，那么就是true，否则就是false
	// First loop, if the robot returns to the origin, it is true, otherwise it is false
	for _, v := range instructions {
		switch v {
		case 'G':
			robot.Go()
		case 'L':
			robot.TurnLeft()
		case 'R':
			robot.TurnRight()
		}
	}

	// ！！！注意：如果遍历完指令后，朝向变了，那么最终一定会回到原点！！！
	// 如果一次循环结束后，机器人的方向不是北方，那么就需要重复执行指令，直到机器人的方向是北方
	// If the direction of the robot is not north after the first loop, the instruction needs to be repeated until the direction of the robot is north
	for robot.facing != "north" {
		for _, v := range instructions {
			switch v {
			case 'G':
				robot.Go()
			case 'L':
				robot.TurnLeft()
			case 'R':
				robot.TurnRight()
			}
		}
	}

	// 当循环结束且机器人朝向北方，如果机器人回到原点，那么就是true，否则就是false
	// When the loop ends and the robot faces north, if the robot returns to the origin, it is true, otherwise it is false
	if robot.position == [2]int{0, 0} {
		return true
	} else {
		return false
	}

}

// 机器人结构定义
// Robot struct definition
type Robot struct {
	facing   string // 机器人朝向 Robot facing
	position [2]int // 机器人位置 Robot position
}

// 机器人构造函数
// Robot constructor
func NewRobot() *Robot {
	return &Robot{
		facing:   "north",      // 机器人默认朝向北方 The robot defaults to the north
		position: [2]int{0, 0}, // 机器人默认位置 The robot defaults to the origin
	}
}

// 机器人Go方法
// Robot Go method
func (robot *Robot) Go() {
	switch robot.facing {
	case "north":
		robot.position[1] += 1
	case "east":
		robot.position[0] += 1
	case "south":
		robot.position[1] -= 1
	case "west":
		robot.position[0] -= 1
	}
}

// 机器人左转方法
// Robot left turn method
func (robot *Robot) TurnLeft() {
	switch robot.facing {
	case "north":
		robot.facing = "west"
	case "east":
		robot.facing = "north"
	case "south":
		robot.facing = "east"
	case "west":
		robot.facing = "south"
	}
}

// 机器人右转方法
// Robot right turn method
func (robot *Robot) TurnRight() {
	switch robot.facing {
	case "north":
		robot.facing = "east"
	case "east":
		robot.facing = "south"
	case "south":
		robot.facing = "west"
	case "west":
		robot.facing = "north"
	}
}

// ==============================================================================================================================
// 2. shorter solution from others in leetcode
func isRobotBounded1(instructions string) bool {
	// x, y 代表机器人的坐标，degree代表机器人的方向
	x, y, degree := 0, 0, 0

	// 遍历指令
	for _, i := range instructions {
		// 如果是左转，degree+3，如果是右转，degree+1，然后都continue，进入下一次循环
		if i == 'R' {
			degree = (degree + 1) % 4
			continue
		} else if i == 'L' {
			degree = (degree + 3) % 4
			continue
		}

		// 如果是前进，根据degree的值，更新x，y的值
		switch degree {
		case 0:
			y++
		case 1:
			x++
		case 2:
			y--
		case 3:
			x--
		}
	}

	// 如果机器人回到原点，或者degree不等于0，那么就是true，否则就是false
	// ！！！注意：如果遍历完指令后，朝向变了，那么最终一定会回到原点！！！
	return x == 0 && y == 0 || degree != 0
}
