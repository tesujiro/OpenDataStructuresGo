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
 2.5 DualArrayDeque: Building a Deque from Two Stacks
 2.6 RootishArrayStack: A Space-Efficient Array Stack

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
 5.1 ChainedHashTable: Hashing with Chaining  
  -> Done  
 5.2 LinearHashTable: Linear Probing  
  -> Done  

6. Binary Trees  
 6.2 BinarySearchTree: An Unbalanced Binary Search Tree  
  -> Done  

7. Random Binary Search Trees  
 7.2 Treap: A Randomized Binary Search Tree  
  -> Done  

8. Scapegoat Trees  
  -> Not Complete  

9. Red-Black Trees  

10. Heaps  

11. Sorting Algorithms  
 11.1 Comparison-Based Sorting  
  11.1.1 Merge-Sort  
   -> Done  
  11.1.2 Quicksort  
   -> Done  
  11.1.3 Heap-sort  
 11.2 Counting Sort and Radix Sort  
  11.2.1 Counting Sort  
   -> Done  
  11.2.2 Radix-Sort  
   -> Done  

12. Graphs  

13. Data Structures for Integers  
 13.1 BinaryTrie: A digital search tree  
  -> Done  
 13.2 XFastTrie: Searching in Doubly-Logarithmic Time  
 13.3 YFastTrie: A Doubly-Logarithmic Time SSet  

14. External Memory Searching  
　　　
## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/tesujiro/OpenDataStructuresGo
BenchmarkList_Ch02_ArrayStack/AddFirst-4 		  200000	    173495 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayStack/AddLast-4 	 	20000000	       150 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayStack/AddRandom-4         	  300000	    127008 ns/op	      63 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque/AddFirst-4          	10000000	       173 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque/AddLast-4           	10000000	       153 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque/AddRandom-4         	  100000	    255583 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_Ch02_DualArrayDeque/AddFirst-4      	   10000	    118224 ns/op	  166969 B/op	       3 allocs/op
BenchmarkList_Ch02_DualArrayDeque/AddLast-4       	   10000	    110730 ns/op	  166968 B/op	       2 allocs/op
BenchmarkList_Ch02_DualArrayDeque/AddRandom-4     	   10000	    117014 ns/op	  166968 B/op	       3 allocs/op
BenchmarkList_Ch02_RootishArrayStack/AddFirst-4   	   20000	    174129 ns/op	      57 B/op	       2 allocs/op
BenchmarkList_Ch02_RootishArrayStack/AddLast-4    	10000000	       109 ns/op	      57 B/op	       2 allocs/op
BenchmarkList_Ch02_RootishArrayStack/AddRandom-4  	   30000	    133225 ns/op	      57 B/op	       2 allocs/op
BenchmarkList_Ch03_DLList/AddFirst-4        	 	10000000	       131 ns/op	      40 B/op	       2 allocs/op
BenchmarkList_Ch03_DLList/AddLast-4         	 	20000000	       132 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_Ch03_DLList/AddRandom-4       	 	  200000	    978381 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_Ch03_SEList/AddFirst-4        	 	 2000000	       645 ns/op	     271 B/op	       8 allocs/op
BenchmarkList_Ch03_SEList/AddLast-4         	 	20000000	       112 ns/op	      56 B/op	       1 allocs/op
BenchmarkList_Ch03_SEList/AddRandom-4       	 	 1000000	     71403 ns/op	     269 B/op	       8 allocs/op
enchmarkList_Ch04_SkiplistList/AddFirst-4         	 5000000	       411 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_Ch04_SkiplistList/AddLast-4          	 3000000	       482 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_Ch04_SkiplistList/AddRandom-4        	 1000000	      2032 ns/op	     122 B/op	       4 allocs/op

BenchmarkSSet_Ch04_SkiplistSSet/AddFirst-4         	 2000000	       755 ns/op	     233 B/op	       4 allocs/op
BenchmarkSSet_Ch04_SkiplistSSet/AddRandom-4        	 1000000	      3801 ns/op	     228 B/op	       4 allocs/op
BenchmarkSSet_Ch04_SkiplistSSet/FindFrom1M-4       	  500000	      3994 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable/AddFirst-4         	 1000000	      1279 ns/op	     160 B/op	       5 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable/AddRandom-4        	 1000000	      1292 ns/op	     158 B/op	       5 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable/FindFrom1M-4       	 1000000	      1078 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch05_LinearHashTable/AddFirst-4          	 5000000	       491 ns/op	     122 B/op	       1 allocs/op
BenchmarkSSet_Ch05_LinearHashTable/AddRandom-4         	 5000000	       433 ns/op	     122 B/op	       1 allocs/op
BenchmarkSSet_Ch05_LinearHashTable/FindFrom1M-4        	 3000000	       413 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch06_BinaryTree/AddFirst-4               	   30000	    107991 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch06_BinaryTree/AddRandom-4              	 1000000	      1493 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch06_BinaryTree/FindFrom1M-4             	 1000000	      1552 ns/op	       8 B/op	       1 allocs/op
BenchmarkSSet_Ch07_Treap/AddFirst-4           		 5000000	       284 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch07_Treap/AddRandom-4          		 1000000	      2462 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch07_Treap/FindFrom1M-4         		  500000	      2323 ns/op	       8 B/op	       1 allocs/op
BenchmarkSSet_Ch13_BinaryTrie/AddFirst-4         	 2000000	       751 ns/op	     104 B/op	       3 allocs/op
BenchmarkSSet_Ch13_BinaryTrie/AddRandom-4        	  300000	      6726 ns/op	    2212 B/op	      46 allocs/op
BenchmarkSSet_Ch13_BinaryTrie/FindFrom1M-4       	 2000000	       628 ns/op	       8 B/op	       1 allocs/op
BenchmarkSSet_Ch13_XFastTrie/AddFirst-4          	 1000000	      1485 ns/op	     294 B/op	       6 allocs/op
BenchmarkSSet_Ch13_XFastTrie/AddRandom-4         	   50000	     29557 ns/op	    7495 B/op	      99 allocs/op
BenchmarkSSet_Ch13_XFastTrie/FindFrom1M-4        	 1000000	      1269 ns/op	     120 B/op	       7 allocs/op

BenchmarkSort_Ch11_MergeSort/NoSort-4         	 	3000000		       565 ns/op	     352 B/op	       1 allocs/op
BenchmarkSort_Ch11_MergeSort/Reverse-4        	 	3000000		       490 ns/op	     352 B/op	       1 allocs/op
BenchmarkSort_Ch11_MergeSort/Random-4         	 	2000000		       896 ns/op	     342 B/op	       1 allocs/op
BenchmarkSort_Ch11_QuickSort/NoSort-4         	 	5000000		       320 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort_Ch11_QuickSort/Reverse-4        	 	5000000		       323 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort_Ch11_QuickSort/Random-4         	 	3000000		       702 ns/op	       0 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort/NoSort-4         	200000000	        21.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort/Reverse-4        	200000000	         9.27 ns/op	      13 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort/Random-4         	10000000	       178 ns/op	     115 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort/NoSort-4            	30000000	        40.8 ns/op	      24 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort/Reverse-4           	30000000	        48.2 ns/op	      32 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort/Random-4            	50000000	        27.0 ns/op	      32 B/op	       0 allocs/op
```
