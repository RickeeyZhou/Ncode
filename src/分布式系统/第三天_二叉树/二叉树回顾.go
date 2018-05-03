package main

import (
	"github.com/cheekybits/genny/generic"
	"sync"
	"fmt"
)
type item generic.Type
//定义二叉树的类型,性质,
//节点,二叉树
type  node struct {

	key int
	value item
	left *node
	right *node
}
type tree struct {
	root *node
	lock sync.RWMutex // 读写锁,一定会对二叉树进行读写操作
}

func main() {
	var onetree tree
	fmt.Println(onetree)
	//onetree.(8,8)
	onetree.insert(8, 8) //为啥是
	a := &onetree
	a.insert(8,8) //
	onetree.insert(9,9)
	(&onetree).insert(8,8)
	(&onetree).insert(8,8)
	(&onetree).insert(8,8)
	(&onetree).insert(8,8)
	(&onetree).insert(8,8)
	(&onetree).insert(8,8)
	onetree.insert(5,5)
	onetree.insert(7,7)
	onetree.insert(2,9)
	onetree.stringPrint()
	bool:=onetree.Search(100)
	fmt.Println(bool)
	fmt.Println(*onetree.max1())

}
//}
//二叉树的插入
func(onetree *tree)insert(key int,value item){
onetree.lock.Lock()
defer onetree.lock.Unlock()

newnode:=node{key,value,nil,nil}
if onetree.root==nil{
	onetree.root=&newnode        // 创建根节点,这里的根节点需要地址
}else{
	insertNode1(onetree.root,&newnode)
}
}
func insertNode1(node,newnode *node){
	// 判断左还是右
	// 然后如果往左非空 ,又要判断是新值和旧值大小,并且进入递归调用,一直找到
	if node.key>newnode.key{
		if node.left==nil{
			node.left=newnode
		}else{
			insertNode1(node.left,newnode)
		}
	}else{
		if node.right==nil{
			node.right=newnode
		}else{
			insertNode1(node.right,newnode)
		}
	}
}//  (8,8) (5,5)   插入(7)会怎么样 ?输出到 第三层5,5 的右边,并不会发生,将原来数据挤走的情况.
//二叉树的查找
func (tree *tree)Search(key int)bool{
	tree.lock.Lock()
	defer tree.lock.Unlock()

	return search(tree.root,key)
}
func search( node *node,key int)bool{
	if node==nil{
		return  false
	}
	if node.key >key{
		 return search(node.left,key)
	}
	if node.key <key{
		return search(node.right,key)
	}
	return true
}
// 最大最小
func (tree *tree)max1()*item{ // 最大k 即最右节点
	tree.lock.Lock()
	defer tree.lock.Unlock()

	node:=tree.root
	if node==nil{
		return nil
	}
	for {
		if node.right==nil{
			return &node.value
		}
		node=node.right
	}

}
// 前中后的遍历
//删除
func (tree *tree)Remove(key int){
	tree.lock.Lock()
	defer tree.lock.Unlock()
	remove(tree.root,key)
}
func remove(node *node,key int){

}
//格式化输出
func (tree *tree)stringPrint(){
	tree.lock.Lock()
	defer (*tree).lock.Unlock()
	fmt.Println("-----------------------------")
	if tree.root==nil{
		fmt.Println("this tree is empty")
		return
	}
	stringify(tree.root,0)
	fmt.Println("-----------------------------")
}
func stringify(node *node,level int){
	if node==nil{
		return
	}
		format:=""

		for i:=0;i<level;i++{
			format+="\t"
		}
		format+="-----["
		level++
		stringify(node.left,level)
		fmt.Println(format,node.key)
		stringify(node.right,level)
}