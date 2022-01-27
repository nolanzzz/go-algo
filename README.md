# GoAlgo
本仓库包含用Golang解决的Leetcode算法题题解以及笔记，学习路线主要跟随[代码随想录](https://github.com/youngyangyang04/leetcode-master)

题目以数据结构或算法类型来分类，目录如下：

## 数组
- [26.删除有序数组中的重复项](Array/26删除有序数组中的重复项.md)
- [27.移除元素](Array/27移除元素.md)
- [34.在排序数组中查找元素的第一个和最后一个位置](Array/34在排序数组中查找元素的第一个和最后一个位置.md)
- [35.搜索插入位置](Array/35搜索插入位置.md)
- [51.螺旋矩阵](Array/51螺旋矩阵.md)
- [59.螺旋矩阵II](Array/59螺旋矩阵II.md)
- [69.x的平方根](Array/69x的平方根.md)
- [76.最小覆盖子串](Array/76最小覆盖子串.md)
- [209.长度最小的子数组](Array/209长度最小的子数组.md)
- [283.移动零](Array/283移动零.md)
- [367.有效的完全平方数](Array/367有效的完全平方数.md)
- [704.二分查找](Array/704二分查找.md)
- [844.比较含退格的字符串](Array/844比较含退格的字符串.md)
- [904.水果成篮](Array/904水果成篮.md)
- [977.有序数组的平方](Array/977有序数组的平方.md)
- [剑指29.顺时针打印矩阵](Array/剑指29顺时针打印矩阵.md)

## 哈希
- [1.两数之和](Hash/1两数之和.md)
- [15.三数之和](Hash/15三数之和.md)
- [18.四数之和](Hash/18四数之和.md)
- [202.快乐数](Hash/202快乐数.md)
- [242.有效的字母异位词](Hash/242有效的字母异位词.md)
- [349.两个数组的交集](Hash/349两个数组的交集.md)
- [383.赎金信](Hash/383赎金信.md)
- [454.四数相加II](Hash/454四数相加II.md)

## 链表
- [19.删除链表的倒数第n个结点](LinkedList/19删除链表的倒数第n个结点.md)
- [24.两两交换链表中的结点](LinkedList/24两两交换链表中的结点.md)
- [142.环形链表II](LinkedList/142环形链表II.md)
- [203.移除链表元素](LinkedList/203移除链表元素.md)
- [206.反转链表](LinkedList/206反转链表.md)
- [707.设计链表](LinkedList/707设计链表.md)
- [面试题0207.链表相交](LinkedList/面试题0207链表相交.md)

## 栈与队列
- [20.有效的括号](StackQueue/20有效的括号.md)
- [150.逆波兰表达式求值](StackQueue/150逆波兰表达式求值.md)
- [225.用队列实现栈](StackQueue/225用队列实现栈.md)
- [232.用栈实现队列](StackQueue/232用栈实现队列.md)
- [239.滑动窗口最大值](StackQueue/239滑动窗口最大值.md)
- [347.前k个高频元素](StackQueue/347前k个高频元素.md)
- [1047.删除字符串中的所有相邻重复项](StackQueue/1047删除字符串中的所有相邻重复项.md)

## 字符串
- [28.实现strStr()](String/28实现strStr().md)
- [151.反转字符串里的单词](String/151反转字符串里的单词.md)
- [344.反转字符串](String/344反转字符串.md)
- [459.重复的子字符串](String/459重复的子字符串.md)
- [541.反转字符串II](String/541反转字符串II.md)
- [剑指05.替换空格](String/剑指05替换空格.md)
- [剑指58II.左旋转字符串](String/剑指58II左旋转字符串.md)

## 二叉树

### 二叉树的理论基础
- [理论基础](Tree/二叉树理论基础.md)

### 二叉树的遍历方式
- DFS-深度优先遍历
  - 迭代法：通过栈来模拟递归
  - [144.二叉树的前序遍历](Tree/DFS/144二叉树的前序遍历.md)
  - [94.二叉树的中序遍历](Tree/DFS/94二叉树的中序遍历.md)
  - [145.二叉树的后序遍历](Tree/DFS/145二叉树的后序遍历.md)
  - [二叉树的统一迭代法](Tree/DFS/二叉树的统一迭代法.md)
- BFS-广度优先遍历
  - [102.二叉树的层序遍历](Tree/BFS/102二叉树的层序遍历.md)：通过队列模拟
  - 通过BFS解决的题目：
  - [107.二叉树的层序遍历II](Tree/BFS/107二叉树的层序遍历II.md)
  - [116.填充每个节点的下一个右侧节点指针](Tree/BFS/116填充每个节点的下一个右侧节点指针.md)
  - [117.填充每个节点的下一个右侧节点指针II](Tree/BFS/117填充每个节点的下一个右侧节点指针II.md)
  - [199.二叉树的右视图](Tree/BFS/199二叉树的右视图.md)
  - [429.N叉树的层序遍历](Tree/BFS/429N叉树的层序遍历.md)
  - [515.在每个树行中找最大值](Tree/BFS/515在每个树行中找最大值.md)
  - [559.N叉树的最大深度](Tree/BFS/559N叉树的最大深度.md)
  - [637.二叉树的层平均值](Tree/BFS/637二叉树的层平均值.md)

### 求二叉树的属性
- [101.二叉树：是否对称](Tree/101对称二叉树.md)，[100.二叉树：是否相同](Tree/100相同的树.md)
  - 递归：后序，比较的是根节点的左子树和右子树是不是相互反转
  - 迭代：使用栈/队列将两个节点顺序放入容器中进行比较
- [572.二叉树：是否为另一棵树的子树](Tree/572另一棵树的子树.md)
  - 递归：查看是否是相同的树，递归
- [104.二叉树：求最大深度](Tree/BFS/104二叉树的最大深度.md)
  - 递归：后序，求根节点最大高度就是最大深度，通过递归函数的返回值计算树的高度（左树和右树的最大高度+1）
  - 迭代：层序，层数就是最大深度
- [111.二叉树：求最小深度](Tree/BFS/111二叉树的最小深度.md)
  - 递归：后序，求根节点的最小高度就是最小深度，注意最小深度的定义
  - 迭代：层序
- [222.二叉树：求有多少个节点](Tree/222完全二叉树的节点个数.md)
  - 递归：后序，通过递归函数的返回值计算节点数量
  - 递归2：如果是完全二叉树，左右指针分别遍历至最后一层，利用特性来计算
  - 迭代：层序
- [110.二叉树：是否平衡](Tree/110平衡二叉树.md)
  - 递归：后序，先求左右子树的高度，差值大于1时直接返回-1
- [257.二叉树：找所有路径](Tree/257二叉树的所有路径.md)
  - 递归：前序，path和res数组分别记录目前path和最终结果。注意回溯的思想
- [404.二叉树：求左叶子之和](Tree/404左叶子之和.md)
  - 递归：后序，注意判断左叶子的条件
  - 迭代：直接用栈模拟后序
- [513.二叉树：找左下角的值](Tree/513找树左下角的值.md)
  - 递归：需要记录最大深度，来判断是否为最后一层
  - 迭代：层序遍历找到最后一行的第一个值
- [112.二叉树：求路径总和](Tree/112路径总和.md)，[路径总和II](Tree/113路径总和II.md)
  - 递归：顺序无所谓，递归函数返回值为bool是为了搜索一边，没有返回值是搜索整棵树

### 二叉树的修改与改造
- [226.反转二叉树](Tree/226反转二叉树.md)
  - 递归：前序，交换左右孩子
  - 迭代：前序或层序都行
- [106.中序与后序构造二叉树](Tree/106从中序与后序遍历序列构造二叉树.md)，[105.前序与中序构造二叉树](Tree/105从前序与中序遍历序列构造二叉树.md)
  - 递归：前序，以后序遍历序列的最后一个节点为root，寻找前/中序和后序中的切割点分别构造左右区间
- [654.构造最大的二叉树](Tree/654最大二叉树.md)
  - 递归：前序，找到数组的最大值作为切割点，构造左右区间
- [617.合并两个二叉树](Tree/617合并二叉树.md)
  - 递归：前序，同时操作两个节点，可以直接用其中一个节点而不用新建root
  - 迭代：层序，每次往队列中放入两个节点

### 求二叉搜索树的属性
- BST的遍历不需要额外存储空间，因为无需回溯
- [700.二叉搜索树中的搜索](Tree/700二叉搜索树中的搜索.md)
  - 递归：按大小判断递归
  - 迭代：按大小判断遍历
- [98.是不是二叉搜索树](Tree/98验证二叉搜索树.md)
  - 递归1：中序遍历，判断生成的数组是否为一个递增有序数组
  - 递归2：无需数组，递归时更新min和max值来判断
- [530.求二叉搜索树的最小绝对差](Tree/530二叉搜索树的最小绝对值.md)
  - 递归：中序，用pre来记录前一个结点方便计算差值，双指针
  - 迭代：中序，相同逻辑
- [501.求二叉搜索树中的众数](Tree/501二叉搜索树中的众数.md)
  - 递归：中序，count和maxCount,更新maxCount时清空已存在的res
  - 迭代：中序，相同逻辑
- [538.把二叉搜索树转换为累加树](Tree/538把二叉搜索树转换为累加树.md)
  - 递归：中序（右中左），双指针pre来操作累加
  - 迭代：栈模拟中序，相同逻辑

### 二叉树公共祖先问题
- [236.二叉树的最近公共祖先](Tree/236二叉树的最近公共祖先.md)
  - 递归：后序，然后判断几种情况
- [235.二叉搜索树的最近公共祖先](Tree/235二叉搜索树的最近公共祖先.md)
  - 递归：根据root的数值比较决定递归方向
  - 迭代：直接指针遍历，与递归相同逻辑

### 二叉搜索树的修改与改造
- [701.二叉搜索树中的插入操作](Tree/701二叉搜索树中的插入操作.md)
  - 递归：找到null值时插入
  - 迭代：遍历找到位置插入
- [450.二叉搜索树中的删除操作](Tree/450删除二叉搜索树中的节点.md)
  - 递归：前序，重平衡时用右树的最小值作为新的root值
- [669.修剪二叉搜索树](Tree/669修剪二叉搜索树.md)
  - 递归：前序
- [108.构造二叉搜索树](Tree/108将有序数组转换为二叉搜索树.md)
  - 递归：前序，用数组中间节点作为root进行区间分隔

## 回溯算法

### 回溯算法理论基础
- [理论基础](Backtracking/回溯算法理论基础.md)

### 组合问题
- 组合问题其实也是一种子集问题
- [77.组合](Backtracking/77组合.md)
  - 回溯模板，res，path，start递归
- [216.组合总和III](Backtracking/216组合总和III.md)
  - 回溯模板，res，path，start递归
  - 还需要sum记录总和
- [17.电话号码的字母组合](Backtracking/17电话号码的字母组合.md)
  - 回溯模板，res，path，index递归
  - 数组来储存数字与字母的映射
- [39.组合总和](Backtracking/39组合总和.md)
  - 回溯模板，res，path，index递归
  - 由于可重复，回溯时用i而不是i+1
  - 若先排序数组，可以剪枝
- [40.组合总和II](Backtracking/40组合总和II.md)
  - 回溯模板，res，path，index递归
  - 不可重复，回溯时用i+1
  - 需要去重操作，首先要排序数组
  - 然后递归逻辑中每次要去重（candidates[i] == candidates[i-1])
  - 可以剪枝

### 分割问题
- 分隔与组合相似，组合选择变成了切割点
- [131.分割回文串](Backtracking/131分隔回文串.md)
  - 回溯模板，res，path，start递归
  - 判断回文，不符合时continue
  - 终止条件为到达末尾
- [93.复原IP地址](Backtracking/93复原IP地址.md)
  - 回溯模板，res，path，start递归
  - 判断每段地址是否合法
  - 终止条件为1.到达末尾；2.path有四段地址

### 子集问题
- 求取子集问题，不需要任何剪枝！因为子集就是要遍历整棵树
- [78.子集](Backtracking/78子集.md)
  - 回溯模板，res，path，start递归
  - 每次递归先保存当前path至res
  - 终止条件可以省略，因为for循环遍历至最后一个数字时自然结束
- [90.子集II](Backtracking/90子集II.md)
  - 与78.子集一个套路
  - 需要去重操作，首先要排序数组
  - 然后递归逻辑中每次要去重
- [491.递增子序列](Backtracking/491递增子序列.md)
  - 回溯模板，res，path，start递归
  - 根据题目条件，当path有两个元素时即可保存，保存后继续往后走
  - 需用used来记录本层用过的数字，以免重复搜索同样的数字开头的子序列

### 排列问题
- 排列问题是有序的，所以每次循环都要从0开始，不需要start参数
- 去重时需要用used来记录
- [46.全排列](Backtracking/46全排列.md)
  - 回溯模板，res，path
  - used记录出现过的元素，同样要回溯
- [47.全排列II](Backtracking/47全排列II.md)
  - 回溯模板，res，path
  - used记录出现过的元素，同样要回溯
  - 需要去重：
    1. 已经搜索过以第i个数字开头的排列，跳过
    2. 重复的数字，规定按从左到右的顺序搜索，还没有搜索过第i-1个数时跳过

### 重新安排行程（图论额外拓展）
- [332.重新安排行程](Backtracking/332重新安排行程.md)
  - 回溯模板，res，不需要path了，因为此时res相当于path，找到一条符合要求的path时便结束了
  - 需要一个map来记录起点和终点的映射
  - map需要根据终点排序，这样最终的结果才是根据字典排序的
  - 终止条件：当res的路径数等于tickets数加一时结束

### 棋盘问题
- [51.N皇后](Backtracking/51N皇后.md)
  - 回溯模板，board来记录棋盘状态，row记录当前的行
  - 递归遍历行
  - 递归里的for遍历列
  - 终止条件：当row = n时，完成了所有行
  - 检查是否合法时，查看列、45度和135度上斜线即可
- [37.解数独](Backtracking/37解数独.md)
  - 回溯模板，board记录棋盘状态
  - 需要遍历所有行所有列，以及1-9的值，找到当前cell合适的值直接返回true
  - 没有合适的值返回false
  - 检查是否合法时，查看行、列、以及九宫格

## 贪心算法

- [贪心算法理论基础](Greedy/贪心算法理论基础.md)
- [455.分发饼干](Greedy/455分发饼干.md)
  - 排序饼干和孩子胃口
  - 双指针
- [376.摆动序列](Greedy/376摆动序列.md)
  - 无需删除操作，只用记录符合条件的节点数
  - 记录currDiff和preDiff用来比较

## 排序
- [快速排序](Sorting/快速排序.md)

## 其他算法
- [KMP算法](SpecialAlgorithms/kmp.go)