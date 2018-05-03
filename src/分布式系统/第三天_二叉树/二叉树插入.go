//package main
//import (
//	"github.com/cheekybits/genny/generic"
//	"sync"
//	"fmt"
//)
//type item generic.Type
//type Node struct {
//	// 节点结构体
//	key        int
//	value      item
//	left_Node  *Node
//	right_Node *Node
//}
//type onetree struct {
//	// 二叉树的结构体
//	root *Node        // 根节点
//	lock sync.RWMutex // 读写锁
//}
//// 查找 插入 删除
//func main() {
//	var tree onetree
//	fmt.Println(tree)
//	fmt.Printf("%T\t%p", tree, &tree)
//	//插入
//	tree.Insert(8,8)
//	tree.Insert(2,2)
//	tree.Insert(10,10)
//	tree.Insert(132,10)
//	tree.Insert(1,10)
//	tree.Insert(13,10)
//	tree.Insert(17,10)
//	tree.string()
//
//}
//// 插入节点
//func (tree *onetree) Insert(key int, value item) {
//	tree.lock.Lock()
//	defer tree.lock.Unlock()
//	// 创建节点
//	node := &Node{key, value, nil, nil}
//	if tree.root==nil{
//		tree.root=node
//	}	else{
//		// 以某个节点为入口插入
//		insertNode(tree.root,node)
//	}
//}
//func insertNode(node, newNode *Node){
//	// 判断节点插入左边还是右边
//	//如果新节点的key< node.root,差左边
//	if newNode.key<node.key{
//		if node.left_Node==nil {
//			node.left_Node = newNode
//		}else{
//			insertNode(node.left_Node,newNode)
//		}
//	}else{
//		if node.right_Node==nil {
//			node.right_Node = newNode
//		}else{
//			insertNode(node.right_Node,newNode)
//		}
//	}
//}
////格式化 输入输出
//func (tree *onetree)string(){
//	tree.lock.Lock()
//	defer tree.lock.Lock()
//
//	if tree.root==nil{
//		fmt.Println("tree is empty")
//		return
//	}
//	stringify(tree.root,0)
//	fmt.Println("----------------------------")
//}
//func stringify(node *Node,level int){
//	if node==nil{
//		return
//	}
//	format:=""
//	for i:=0;i<level;i++{
//		format+="\t"
//	}
//	format+="----["
//	level++
//	// 先递归打印左子树
//	stringify(node.left_Node,level)
//	fmt.Println(format+"&d\n",node.key)
//	//再一次递归 打印右子树
//	stringify(node.right_Node,level)
//}
//// 返回最小值的地址
//func (tree *onetree)min()*item{
//	tree.lock.Lock()
//	defer tree.lock.Unlock()
//
//	node:=tree.root
//	if node==nil{
//		return nil
//	}
//	for {  /// 最小值一定在最左边
//		if node.left_Node==nil{
//			return &node.value
//		}
//		node=node.left_Node
//	}
//}
//// 返回最大值的地址
//func (tree *onetree)max()*item{
//	tree.lock.Lock()
//	defer tree.lock.Unlock()
//
//	node:=tree.root
//	if node==nil{
//		return nil
//	}
//	for {  /// 最大值一定在最右边
//		if node.right_Node==nil{
//			return &node.value
//		}
//		node=node.right_Node
//	}
//}
//// 搜索, 查询某个key
//func (tree *onetree)Search(key int )bool{
//	//
//	tree .lock.Lock()
//	defer  tree.lock.Unlock()
//	return search(tree.root,key)
//}
//func search(node *Node,key int )bool{
//	if node==nil{
//		return false
//	}
//	if key<node.key{
//		return search(node.left_Node,key)
//	}
//	if key>node.key{
//		return search(node.right_Node,key)
//	}
//	return true
//}
//// 遍历
////func ()