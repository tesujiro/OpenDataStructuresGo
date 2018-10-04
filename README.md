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
BenchmarkList_Ch02_ArrayStack_AddFirst-4       	  200000	    157613 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayStack_AddLast-4        	20000000	       156 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayStack_AddRandom-4      	  300000	    138063 ns/op	      63 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque_AddFirst-4       	10000000	       157 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque_AddLast-4        	10000000	       187 ns/op	      61 B/op	       1 allocs/op
BenchmarkList_Ch02_ArrayDeque_AddRandom-4      	   50000	    140180 ns/op	      49 B/op	       1 allocs/op
BenchmarkList_Ch03_DLList_AddFirst-4           	10000000	       127 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_Ch03_DLList_AddLast-4            	10000000	       115 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_Ch03_DLList_AddRandom-4          	  200000	   1076100 ns/op	      40 B/op	       1 allocs/op
BenchmarkList_Ch03_SEList_AddFirst-4           	 2000000	       744 ns/op	     271 B/op	       8 allocs/op
BenchmarkList_Ch03_SEList_AddLast-4            	10000000	       135 ns/op	      56 B/op	       1 allocs/op
BenchmarkList_Ch03_SEList_AddRandom-4          	 1000000	     79783 ns/op	     269 B/op	       8 allocs/op
BenchmarkList_Ch04_SkiplistList_AddFirst-4    	 3000000	       370 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_Ch04_SkiplistList_AddLast-4     	 5000000	       501 ns/op	     122 B/op	       4 allocs/op
BenchmarkList_Ch04_SkiplistList_AddRandom-4   	 1000000	      2294 ns/op	     122 B/op	       4 allocs/op

BenchmarkSSet_Ch04_SkiplistSSet_AddFirst-4         	 2000000	       783 ns/op	     241 B/op	       4 allocs/op
BenchmarkSSet_Ch04_SkiplistSSet_AddRandom-4        	 1000000	      3498 ns/op	     219 B/op	       4 allocs/op
BenchmarkSSet_Ch04_SkiplistSSet_FindFrom1M-4       	  500000	      2978 ns/op	       8 B/op	       1 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable_AddFirst-4     	 1000000	      1356 ns/op	     159 B/op	       5 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable_AddRandom-4    	 1000000	      1115 ns/op	     158 B/op	       5 allocs/op
BenchmarkSSet_Ch05_ChainedHashTable_FindFrom1M-4   	 2000000	      1023 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch05_LinearHashTable_AddFirst-4      	10000000	       408 ns/op	     122 B/op	       1 allocs/op
BenchmarkSSet_Ch05_LinearHashTable_AddRandom-4     	 5000000	       424 ns/op	     122 B/op	       1 allocs/op
BenchmarkSSet_Ch05_LinearHashTable_FindFrom1M-4    	 5000000	       235 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch06_BinaryTree_AddFirst-4           	   50000	    170023 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch06_BinaryTree_AddRandom-4          	 1000000	      1216 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch06_BinaryTree_FindFrom1M-4         	 1000000	      1211 ns/op	       8 B/op	       1 allocs/op
BenchmarkSSet_Ch07_Treap_AddFirst-4     		 5000000	       289 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch07_Treap_AddRandom-4    		 1000000	      1821 ns/op	      56 B/op	       2 allocs/op
BenchmarkSSet_Ch07_Treap_FindFrom1M-4   		  500000	      2470 ns/op	       7 B/op	       0 allocs/op
BenchmarkSSet_Ch13_BinaryTrie_AddFirst-4     	 	 2000000	       753 ns/op	     104 B/op	       3 allocs/op
BenchmarkSSet_Ch13_BinaryTrie_AddRandom-4    	 	  300000	      7447 ns/op	    2212 B/op	      46 allocs/op
BenchmarkSSet_Ch13_BinaryTrie_FindFrom1M-4   	 	 2000000	       804 ns/op	       8 B/op	       0 allocs/op
BenchmarkSSet_Ch13_XFastTrie_AddFirst-4    	 	 1000000	      1815 ns/op	     294 B/op	       6 allocs/op
BenchmarkSSet_Ch13_XFastTrie_AddRandom-4    		   50000	     47182 ns/op	    7495 B/op	      99 allocs/op
BenchmarkSSet_Ch13_XFastTrie_FindFrom1M-4   		 1000000	      1880 ns/op	     119 B/op	       7 allocs/op

BenchmarkSort_Ch11_MergeSort_NoSort-4    	 3000000	       558 ns/op	     352 B/op	       1 allocs/op
BenchmarkSort_Ch11_MergeSort_Reverse-4   	 3000000	       530 ns/op	     352 B/op	       1 allocs/op
BenchmarkSort_Ch11_MergeSort_Random-4    	 2000000	       844 ns/op	     342 B/op	       1 allocs/op
BenchmarkSort_Ch11_QuickSort_NoSort-4    	 5000000	       340 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort_Ch11_QuickSort_Reverse-4   	 5000000	       328 ns/op	       0 B/op	       0 allocs/op
BenchmarkSort_Ch11_QuickSort_Random-4    	 3000000	       662 ns/op	       0 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort_NoSort-4    	50000000	        20.4 ns/op	      29 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort_Reverse-4   	100000000	        11.2 ns/op	      18 B/op	       0 allocs/op
BenchmarkCSort_Ch11_CountingSort_Random-4    	10000000	       181 ns/op	     115 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort_NoSort-4       	30000000	        47.3 ns/op	      32 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort_Reverse-4      	30000000	        46.4 ns/op	      32 B/op	       0 allocs/op
BenchmarkCSort_Ch11_RadixSort_Random-4       	50000000	        27.1 ns/op	      32 B/op	       0 allocs/op

```
