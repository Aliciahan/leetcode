package data_structure

/*
https://dreamerjonson.com/2019/03/03/golang-7-structure-tree/index.html
通道遍历 特别有趣
 */

type NodeTree struct {
	Val interface{}
	Children []*NodeTree
}

/*
return true when no child
 */
func (nt NodeTree) IsVacant () bool {
	if len(nt.Children) == 0 {
		return true
	}
	return false
}

func (nt NodeTree) NtAppendTo(interface{},interface{}){

}