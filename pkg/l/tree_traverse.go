/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-01 17:26:47
 */
package l

type Stack struct {
	data []*TreeNode
}

func NewStack() Stack {
	return Stack{
		data: []*TreeNode{},
	}
}

func (this *Stack) len() int {

	return len(this.data)
}
func (this *Stack) push(node *TreeNode) {
	if node == nil {
		return
	}
	this.data = append(this.data, node)
}

func (this *Stack) pop() *TreeNode {

	res := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return res
}

func (this *Stack) top() *TreeNode {

	return this.data[len(this.data)-1]
}

func PostorderTraversal(root *TreeNode) []int {

	res := []int{}

	stack := NewStack()

	cur := root
	var pre *TreeNode = nil
	for stack.len() != 0 || cur != nil {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}

		cur = stack.pop()
		if cur.Right == nil || pre == cur.Right {
			res = append(res, cur.Val)
			pre = cur
			cur = nil

		} else {
			stack.push(cur)
			cur = cur.Right
		}

	}

	return res
}

func preorderTraversalNorecur(root *TreeNode) []int {
	res := []int{}
	stack := NewStack()

	stack.push(root)

	for stack.len() != 0 {
		node := stack.pop()
		if node == nil {
			break
		}
		res = append(res, node.Val)
		stack.push(node.Right)
		stack.push(node.Left)
	}
	return res
}

func inorderTraversalNorecur(root *TreeNode) []int {
	res := []int{}

	stack := NewStack()

	cur := root
	for stack.len() != 0 {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}

		cur = stack.pop()
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func LevelTra(root *TreeNode, level int) []int {
	res := []int{}
	cache := []*TreeNode{}
	left := 0
	cache = append(cache, root)
	right := 1
	levelNum := 0
	for left < right {
		levelNum++
		if levelNum > level {
			return res
		}
		for left < right {
			if levelNum == level {
				res = append(res, (*cache[left]).Val)
			}
			if cache[left].Left != nil {
				cache = append(cache, (*cache[left]).Left)
			}
			if cache[left].Right != nil {
				cache = append(cache, (*cache[left]).Right)
			}

			left++
		}
		right = len(cache)
	}
	return res
}

func LevelAverage(root *TreeNode) []float64 {
	res := []float64{}
	cache := []*TreeNode{}
	left := 0
	cache = append(cache, root)
	right := 1

	for left < right {
		levelSum := 0.0
		cnt := 0.0
		for left < right {
			levelSum = levelSum + float64((*cache[left]).Val)
			cnt++
			if (*cache[left]).Left != nil {
				cache = append(cache, (*cache[left]).Left)
			}
			if (*cache[left]).Right != nil {
				cache = append(cache, (*cache[left]).Right)
			}
			left++
		}
		res = append(res, levelSum/cnt)
		right = len(cache)
	}

	return res
}
