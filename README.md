# Open Data Structures in Go

 Study notes of Open Data Structures by Pat Morin in Go.

## Contents

1. Introduction  
 1.2 Interfaces  
  1.2.1 The Queue, Stack, and Deque Interfaces  
     -> Done  
  1.2.2 The List Interface: Linear Sequences  
     -> Done  
  1.2.3 The USet Interface: Unordered Sets  
     -> Done  
  1.2.4 The SSet Interface: Sorted Sets  
     -> Done  

2. Array-Based Lists  
 2.1 ArrayStack: Fast Stack Operations Using an Array  
  -> Done  
 2.2    
 2.3 ArrayQueue: An Array-Based Queue  
  -> Done  
 2.4 ArrayDeque: Fast Deque Operations Using an Array  
  -> Done  
 2.5  　
 2.6  　

3. Linked Lists  
 3.1 SLList: A Singly-Linked List  
  -> Done  
 3.2 DLList: A Doubly-Linked List  
  -> Done  
 3.3 SEList: A Space-Efficient Linked List
  -> Done  

4. Skiplists  
 4.2 SkiplistSSet: An Efficient SSet  
  -> Done  
 4.3 SkiplistList: An Efficient Random-Access List  
  -> Done  

5. Hash Tables  
6. Binary Trees  
 6.2 BinarySearchTree: An Unbalanced Binary Search Tree  
  -> Done  

7. Random Binary Search Trees  
8. Scapegoat Trees  
9. Red-Black Trees  
10. Heaps  
11. Sorting Algorithms  
12. Graphs  
13. Data Structures for Integers  
14. External Memory Searching  
　　　
## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/tesujiro/OpenDataStructuresGo
BenchmarkList_ArrayStack_AddFirst-4       	  200000	    157613 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_ArrayStack_AddLast-4        	20000000	       156 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_ArrayStack_AddRandom-4      	  300000	    138063 ns/op	      63 B/op	       1 allocs/op
BenchmarkList_ArrayDeque_AddFirst-4       	10000000	       157 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_ArrayDeque_AddLast-4        	10000000	       187 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_ArrayDeque_AddRandom-4      	   50000	    140180 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_DLList_AddFirst-4           	10000000	       127 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_DLList_AddLast-4            	10000000	       115 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_DLList_AddRandom-4          	  200000	   1076100 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_SEList_AddFirst-4           	 2000000	       744 ns/op	     271 B/op	       8 allocs/op
BenchmarkList_SEList_AddLast-4            	10000000	       135 ns/op	      56 B/op	       1 allocs/op
BenchmarkList_SEList_AddRandom-4          	 1000000	     79783 ns/op	     269 B/op	       8 allocs/op
BenchmarkList_SkiplistList_AddFirst-4    	 3000000	       370 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_SkiplistList_AddLast-4     	 5000000	       501 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_SkiplistList_AddRandom-4   	 1000000	      2294 ns/op	     122 B/op	       4 allocs/op
BenchmarkSSet_SkiplistSSet_AddRandom-4    	 1000000	      3448 ns/op	     230 B/op	       4 allocs/op
BenchmarkSSet_SkiplistSSet_FindFrom1M-4   	  500000	      3102 ns/op	       8 B/op	       0 allocs/op
BenchmarkSSet_BinaryTree_AddRandom-4      	 1000000	      1312 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_BinaryTree_FindFrom1M-4     	 1000000	      1348 ns/op	       8 B/op	       0 allocs/op
```
