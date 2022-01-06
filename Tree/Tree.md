## Tree (BinaryTree)
- Represent hierarchical data
- Databases (indexing)
- Autocompletion
- Compilers
- Compression (JPEG, MP3)

**完美二叉树/满二叉树-Perfect**： 结点全满，`2^(k+1)-1`个结点

**完全二叉树-Complete**: 除最后一层外为完美二叉树，最后一层叶结点靠左对齐

**完满二叉树-Full/Strict**: 所有结点有0或2个子结点

## 二叉树的实现
- 链式
- 数组(`i, 2*i+1, 2*i+2`)

### Tree Traversals

- **Breadth First / Level Order**
  - 层次遍历/迭代
- **Depth First**
  - Pre-order前序遍历: _Root_, Left, Right
  - In-order中序遍历: Left, _Root_, Right (ascending order)
  - Post-order后序遍历: Left, Right, _Root_(leaves first, then root)

**Height:** bottom-up

 **Depth:** top-down

## 递归函数什么时候需要返回值？什么时候不需要返回值？这里总结如下三点：

- 如果需要搜索整颗二叉树且不用处理递归返回值，递归函数就不要返回值。（这种情况就是113.路径总和ii）
- 如果需要搜索整颗二叉树且需要处理递归返回值，递归函数就需要返回值。 （这种情况我们在236.二叉树的最近公共祖先 中介绍）
- 如果要搜索其中一条符合条件的路径，那么递归一定需要返回值，因为遇到符合条件的路径了就要及时返回。（本题的情况）