package main

import "fmt"

// 单向链表 头节点添加 尾节点添加 中间插入
// 查找指定val 查找指定位置 删除指定值 删除指定位置
// 遍历所有节点 长度
type Node struct {
    Val int
    Next *Node 
}

type LinkList struct {
    HeadNode *Node
}

func (l *LinkList)AddHead(v int) *Node{
    newHead := &Node{Val:v}
    newHead.Next = l.HeadNode
    l.HeadNode = newHead
    return newHead
}

func (l *LinkList)Append(v int) *Node{
    node := &Node{Val:v}
    if l.HeadNode == nil {
        l.HeadNode = node
    }else {
        cur := l.HeadNode
        for cur.Next != nil {
            cur = cur.Next
        }
        cur.Next = node
    }
    return node
}

func (l *LinkList)Length() int {
    count := 0
    cur := l.HeadNode
    for cur != nil {
        cur = cur.Next
        count++
    }
    return count
}

func (l *LinkList)Insert(v int, index int) *Node{
    node := &Node{Val:v}
    if index <= 0 {
        return l.AddHead(v)
    }else if index > l.Length() {
        return l.Append(v)
    }

    cur := l.HeadNode
    for count := 0;count < index - 1;count++ {
        cur = cur.Next
    }
    node.Next = cur.Next
    cur.Next = node

    return node
}

func (l *LinkList)GetVal(v int) *Node {
    cur := l.HeadNode
    for cur != nil {
        if cur.Val == v {
            return cur
        }
        cur = cur.Next
    }
    return nil
}

func (l *LinkList)GetIndex(index int) *Node {
    cur := l.HeadNode
    for i:=0;i < index;i++ {
        if cur == nil {
            return cur
        }
        cur = cur.Next
    }
    return cur
}

func (l *LinkList)DelVal(v int) *Node {
    cur := l.HeadNode
    for cur.Next != nil {
        if cur.Next.Val == v {
            cur.Next = cur.Next.Next
        }else {
            cur = cur.Next
        }
    }
    return nil
}

func (l *LinkList)DelIndex(index int) *Node {
    cur := l.HeadNode
    for i:=0;i < index;i++ {
        if cur == nil {
            return cur
        }
        cur = cur.Next
    }
    return cur
}

func (l *LinkList)ShowList(){
    cur := l.HeadNode
    for cur != nil {
        fmt.Println(cur.Val)
        cur = cur.Next
    }
}

// 双向链表
type DuNode struct {
    Val int
    Pre *Node
    Next *Node
}

func main() {
    list := LinkList{HeadNode:&Node{Val:1}}

    list.Append(2)
    list.Append(3)
    list.Append(4)
    list.Append(5)
    list.Insert(10, 6)

    fmt.Println(list.GetVal(10))
    fmt.Println(list.GetIndex(2))

    fmt.Println(list.DelVal(5))

    list.ShowList()
}
