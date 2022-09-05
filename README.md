# AC 自动机

本 AC 自动机用于在给定文本中, 找到所有匹配的、预先给定的词条

# 使用入门

initialize

```go
package main

var entries []string
var matcher = NewAC(entries)
```

match text

```go
package main

var text string
var matches = matcher.Match(text)
```

# 实现原理

1. 兼容中文 使用 rune 作为单元, []rune 作为词条
2. 处理文本为 []string

> 2.1. 下标 idx 作为命中后的返回值  
> 2.2. 返回值为 []int, 通过 append 添加所有命中结果

3. ac 构建

> 3.1. 构建 trie tree, 所有叶子节点均存储对应文本的 idx  
> 3.2. 构建 fail tree
>> 3.2.1. root 的 children 的 fail 都指向 root  
> > 3.2.2. 宽度遍历 trie tree, 当前 node 的 child 去匹配所有的 node.fail.children, node.fail...fail.children, 如果两个 child 相等, 则 nod.child.fail 指向匹配的 child  
> > 3.2.3. 如果匹配不到, fail = root

4. 匹配原则

> 4.1. 从 root 开始找, 命中第一个目标 A, 加入 output  
> 4.2. 继续找 A的 []child, 如果继续命中目标， 加入 output  
> 4.3. 如果所有 child 都没命中, 跳转到 fail (包括 fail 到 root)  
