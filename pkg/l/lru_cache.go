package main

func main() {
    lru := Constructor(2)
    lru.Put(1,1)
    lru.Put(2,2)
    lru.Get(1)
    lru.Put(3,3)
    lru.Get(2)
    lru.Put(4,4)
    lru.Get(1)
    lru.Get(3)
    lru.Get(4)
}
// 搞引用一定要用 指针
type LRUCache struct {
    cache    map[int]*Node
    capacity int
    head     *Node
    tail     *Node
}

type Node struct {
    key   int
    value int
    pre   *Node
    next  *Node
}

func Constructor(capacity int) LRUCache {
    n1 := &Node{-1,-1,nil,nil,}
    n2 := &Node{-2,-2,nil,nil,}
    n1.next = n2
    n2.pre = n1

    return LRUCache{
        cache:    make(map[int]*Node, capacity),
        capacity: capacity,
        head:     n1,
        tail:     n2,
    }

}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.removeNode(node)
        this.toHead(node)
        return node.value
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        node.value = value
        this.cache[key] = node
        this.removeNode(node)
        this.toHead(node)
    } else {
        node = &Node{
            key:   key,
            value: value,
        }
        this.cache[key] = node
        this.toHead(node)
        if len(this.cache) > this.capacity {
            this.tail = this.tail.pre
            this.tail.next = nil
            delete(this.cache, this.tail.key)
        }
    }

}

func (this *LRUCache) removeNode(node *Node) {
    node.next.pre = node.pre
    node.pre.next = node.next
}

func (this *LRUCache) toHead(node *Node) {
    node.next = this.head.next
    this.head.next.pre = node
    this.head.next = node
    node.pre = this.head
}
