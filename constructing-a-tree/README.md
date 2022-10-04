
# Constructing a tree

- Problem from HackerRank ([link](https://www.hackerrank.com/contests/hack-the-interview-vi-asia-pacific/challenges/constructing-a-tree)

You have been provided with two integers _N_ and _X_ . Your task is to construct an undirected tree of _N_ nodes such that it has score exactly equal to _X_.

The score of a tree is defined to be the sum of distance of each node from the root node. Node  is the root of the tree. If there is no such tree possible print out a single edge containing "-1 -1"

For example, for the following tree:

![tree-1](tree-1.png)

Score would be: `0 + 1 + 1 + 2 + 2 + 2 = 8`, which is the sum of distance of each node.

If there are multiple possible answers, print any of them.

## Input Format

The first line contains a single integer _T_ denoting the number of test cases.

The first and only line of each test case contains two space separated integers, _N_ and _X_ .

## Constraints

- 1 <= _T_ <= 10
- 1 <= _N_ <= 5 * 10<sup>4</sup>
- 1 <= _X_ <= 10<sup>12</sup>
- Sum of _N_ over all the test cases do not exceed 10<sup>5</sup>

## Output Format

For each test case,

If such a tree is possible, print _N - 1_ lines each containing two space separated integers _u_ and _v_ denoting an edge between node _u_ and node _v_. The output edges must form a tree which is connected and have no multi edges or loops.

Else, print a single line containing "-1 -1"

## Sample Input 0

```plain
1
5 7
```

## Sample Output 0

```plain
1 2
1 3
3 4
4 5
```

## Explanation 0

One of the possible answers is:

![tree-2](tree-2.png)

The sum of distances here is  `0 + 1 + 1 + 2 + 3 = 7`

## Sample Input 1

```plain
2
5 6
3 1
```

## Sample Output 1

```plain
1 2
1 3
3 4
3 5
-1 -1
```

## Explanation 1

For the first test case, one of the possible answers is:

![tree-3](tree-3.png)

The sum of distances here is `0 + 1 + 1 + 2 + 2 = 6`.

For the second test case, we can see that there is no such tree possible. Hence we need to print "-1 -1"
