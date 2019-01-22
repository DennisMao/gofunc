package trees

type TreeInterface interface {
	Init()
	Height()     //树高度
	Size()       //树规模
	InsertAsLC() //作为左孩子插入新节点
	InsertAsRC() //作为右孩子插入新节点
	Succ()       //当前节点的直接后继(中序)
	TravLevel()  //子树层次遍历
	TravPre()    //子树先序遍历
	TravIn()     //子树中序遍历
	TravPost()   //子树后序遍历
}
