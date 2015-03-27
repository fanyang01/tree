# Trees

This package implements several balance trees:

- Red-Balck Tree
- Left-Leaning Red-Black Tree
- Treap
- Splay Tree
- AVL Tree

Following is two benchmark results:

> $ ./benchmark.sh 1000000
> avl
> PASS
> BenchmarkInsert    1000000    1726 ns/op
> BenchmarkSearch    1000000    417 ns/op
> BenchmarkDelete    1000000    2429 ns/op
> ok    github.com/fanyang01/tree/avl    39.863s
> llrb
> PASS
> BenchmarkBU23Insert    1000000    842 ns/op
> BenchmarkBU23Search    1000000    405 ns/op
> BenchmarkBU23Delete    1000000    2576 ns/op
> BenchmarkTD234Insert    1000000    915 ns/op
> BenchmarkTD234Search    1000000    411 ns/op
> BenchmarkTD234Delete    1000000    2572 ns/op
> ok    github.com/fanyang01/tree/llrb    77.026s
> rbtree
> PASS
> BenchmarkInsert    1000000    803 ns/op
> BenchmarkSearch    1000000    406 ns/op
> BenchmarkDelete    1000000    1104 ns/op
> ok    github.com/fanyang01/tree/rbtree    34.366s
> splay
> PASS
> BenchmarkInsert    1000000    187 ns/op
> BenchmarkSearch    1000000    613 ns/op
> BenchmarkDelete    1000000    670 ns/op
> ok    github.com/fanyang01/tree/splay    28.949s
> treap
> PASS
> BenchmarkInsert    1000000    520 ns/op
> BenchmarkSearch    1000000    558 ns/op
> BenchmarkDelete    1000000    1199 ns/op
> ok    github.com/fanyang01/tree/treap    27.100s
> 
> $ ./benchmark.sh 50000
> avl
> PASS
> BenchmarkInsert    50000    1291 ns/op
> BenchmarkSearch    50000    397 ns/op
> BenchmarkDelete    50000    1296 ns/op
> ok    github.com/fanyang01/tree/avl    5.831s
> llrb
> PASS
> BenchmarkBU23Insert    50000    668 ns/op
> BenchmarkBU23Search    50000    395 ns/op
> BenchmarkBU23Delete    50000    1549 ns/op
> BenchmarkTD234Insert    50000    738 ns/op
> BenchmarkTD234Search    50000    386 ns/op
> BenchmarkTD234Delete    50000    1619 ns/op
> ok    github.com/fanyang01/tree/llrb    8.769s
> rbtree
> PASS
> BenchmarkInsert    50000    618 ns/op
> BenchmarkSearch    50000    399 ns/op
> BenchmarkDelete    50000    530 ns/op
> ok    github.com/fanyang01/tree/rbtree    3.105s
> splay
> PASS
> BenchmarkInsert    50000    214 ns/op
> BenchmarkSearch    50000    599 ns/op
> BenchmarkDelete    50000    435 ns/op
> ok    github.com/fanyang01/tree/splay    2.363s
> treap
> PASS
> BenchmarkInsert    50000    427 ns/op
> BenchmarkSearch    50000    509 ns/op
> BenchmarkDelete    50000    424 ns/op
> ok    github.com/fanyang01/tree/treap    2.679s
