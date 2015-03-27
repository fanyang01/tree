# Trees

This package implements several balance trees:

- Red-Balck Tree
- Left-Leaning Red-Black Tree
- Treap
- Splay Tree
- AVL Tree

Following is two benchmark results:

Type/Operations   | Size     | time per operation
------------------|----------|-------------------
avl             |               |
BenchmarkInsert |    1000000    | 1726 ns/op
BenchmarkSearch |    1000000    | 417 ns/op
BenchmarkDelete |    1000000    | 2429 ns/op
llrb                |               |
BenchmarkBU23Insert |    1000000    | 842 ns/op
BenchmarkBU23Search |    1000000    | 405 ns/op
BenchmarkBU23Delete |    1000000    | 2576 ns/op
BenchmarkTD234Insert |    1000000    | 915 ns/op
BenchmarkTD234Search |    1000000    | 411 ns/op
BenchmarkTD234Delete |    1000000    | 2572 ns/op
rbtree          |
BenchmarkInsert |    1000000    | 803 ns/op
BenchmarkSearch |    1000000    | 406 ns/op
BenchmarkDelete |    1000000    | 1104 ns/op
splay           |               |
BenchmarkInsert |    1000000    | 187 ns/op
BenchmarkSearch |    1000000    | 613 ns/op
BenchmarkDelete |    1000000    | 670 ns/op
treap           |               |
BenchmarkInsert |    1000000    | 520 ns/op
BenchmarkSearch |    1000000    | 558 ns/op
BenchmarkDelete |    1000000    | 1199 ns/op