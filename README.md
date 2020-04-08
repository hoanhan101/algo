# 116+ Coding Interview Problems with Detailed Solutions

[![Build Status](https://travis-ci.com/hoanhan101/algo.svg?branch=master)](https://travis-ci.com/hoanhan101/algo)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoanhan101/algo)
](https://goreportcard.com/report/github.com/hoanhan101/algo)
![](https://img.shields.io/github/stars/hoanhan101/algo)
![](https://img.shields.io/github/forks/hoanhan101/algo)
[![hackernews](https://img.shields.io/badge/hackernews-300%2B-orange)](https://news.ycombinator.com/item?id=21975463)
[![r/golang](https://img.shields.io/badge/r/golang-200%2B-orange)](https://www.reddit.com/r/golang/comments/el4vut/101_coding_interview_problems_with_detailed/)

> [Join my mailing list to get the latest updates here →](https://tinyletter.com/hoanhan)

> [Ultimate Go study guide →](https://github.com/hoanhan101/ultimate-go)

> [Buy me a coffee!](https://www.buymeacoffee.com/aHjIWu6Ck)

## Motivation

I am building a database of most frequently appeared coding interview problems
that I think are the most valuable and productive to spend time on. For each
one, I am including my thoughts of process on how to approach and solve it,
adding well-documented solutions with test cases, time and space complexity
analysis. My goal is to help you get good algorithms and data structures so
that you can prepare better for your next coding interviews.

These resources that I am looking at are:
- [Interview Cake](https://www.interviewcake.com/)
- [LeetCode](https://leetcode.com/)
- Cracking the Coding Interviews
- [Grokking the Coding Interview: Patterns for Coding Questions](https://www.educative.io/courses/grokking-the-coding-interview)
- Elements of Programming Interviews

If you’re interested in getting updates for this, feel free to join my [mailing 
list here →](https://tinyletter.com/hoanhan)

## Table of Contents 

- **[Interview Cake](https://www.interviewcake.com/)**
  - Array and string manipulation
    - [Merge meeting times](interviewcake/merge_meetings_test.go)
    - [Reverse string in place](interviewcake/reverse_string_test.go)
    - [Reverse words](interviewcake/reverse_word_test.go)
    - [Merge sorted arrays](interviewcake/merge_sorted_arrays_test.go)
  - Hashing and hash tables
    - [In-flight entertainment](interviewcake/inflight_test.go)
    - [Permutation palindrome](interviewcake/permutation_palindrome_test.go)
    - [Word cloud data](interviewcake/word_cloud_test.go)
    - [Top scores](interviewcake/top_scores_test.go)
  - Greedy algorithms
    - [Apple stocks](interviewcake/apple_stocks_test.go)
    - [Highest product of three](interviewcake/highest_product_of_three_test.go)
    - [Product of all other numbers](interviewcake/product_of_others_test.go)
    - [In-place shuffle](interviewcake/inplace_shuffle_test.go)
  - Sorting, searching, and logarithms
    - [Find rotation point](interviewcake/find_rotation_point_test.go)
  - Trees and graphs
    - [Balanced binary tree](interviewcake/balanced_binary_tree_test.go)
    - [Binary search tree checker](interviewcake/binary_search_tree_test.go)
    - [2nd largest item in a binary search tree](interviewcake/2nd_largest_item_bst_test.go)
    - [Graph coloring](interviewcake/graph_coloring_test.go)
  - Dynamic programming and recursion
    - [Recursive string permutations](interviewcake/recursive_string_permutation_test.go)
    - [Compute nth Fibonacci number](interviewcake/fibonacci_number_test.go)
    - [Making change](interviewcake/making_change_test.go)
  - Queues and stacks
    - [Largest stack](interviewcake/largest_stack_test.go)
    - [A queue with two stacks](interviewcake/queue_two_stacks_test.go)
    - [Parenthesis matching](interviewcake/parenthesis_matching_test.go)
    - [Bracket validator](interviewcake/bracket_validator_test.go)
  - Linked lists
    - [Delete node](interviewcake/delete_node_test.go)
    - [Linked list has a cycle](interviewcake/linked_list_cycle_test.go)
    - [Reverse a linked list](interviewcake/reverse_linked_list_test.go)
    - [Kth to last node](interviewcake/kth_to_last_test.go)
  - Bit manipulation
    - [The stolen breakfast drone](interviewcake/stolen_breakfast_drone_test.go)
- **[LeetCode](https://leetcode.com/)**
  - Array/String
    - [Two sum I](leetcode/two_sum_i_test.go)
    - [Two sum II](leetcode/two_sum_ii_test.go)
    - [Valid palindrome](leetcode/valid_palindrome_test.go)
    - [Implement strstr()](leetcode/strstr_test.go)
    - [Reverse words in string](leetcode/reverse_words_string_test.go)
    - [Longest substring without repeating characters](leetcode/longest_substring_test.go)
    - [Missing ranges](leetcode/missing_ranges_test.go)
    - [One edit distance](leetcode/one_edit_distance_test.go)
  - Math
    - [Reverse integer](leetcode/reverse_integer_test.go)
    - [Plus one](leetcode/plus_one_test.go)
    - [Palindrome number](leetcode/palindrome_number_test.go)
  - Linked list
    - [Merge sorted linked list](leetcode/merge_sorted_linked_list_test.go)
    - [Add two numbers](leetcode/add_two_numbers_test.go)
    - [Swap nodes in pairs](leetcode/swap_nodes_in_pairs_test.go)
  - Binary tree
    - [Validate binary search tree](leetcode/valid_bst_test.go)
    - [Maximum depth of binary tree](leetcode/max_depth_binary_tree_test.go)
    - [Minimum depth of binary tree](leetcode/min_depth_binary_tree_test.go)
    - [Balanced binary tree](leetcode/balanced_binary_tree_test.go)
    - [Binary tree maximum path sum](leetcode/binary_tree_max_path_sum_test.go)
  - Bit manipulation
    - [Single number I](interviewcake/stolen_breakfast_drone_test.go)
    - [Single number II](leetcode/single_number_ii_test.go)
  - Misc
    - [Spiral matrix](leetcode/spiral_matrix_test.go)
  - Stack
    - [Min stack](leetcode/min_stack_test.go)
    - [Valid parenthesis](leetcode/valid_parentheses_test.go)
  - Dynamic programming
    - [Climbing stairs](leetcode/climbing_stairs_test.go)
- **Cracking the Coding Interviews**
  - Bit manipulation
    - [Bit operation](ctci/bit_operation_test.go)
    - [Bit insertion](ctci/bit_insertion_test.go)
- **[Grokking the Coding Interview: Patterns for Coding Questions](https://www.educative.io/courses/grokking-the-coding-interview)**
  - Sliding Window
    - [Average of any contiguous subarray of size k](gtci/avg_subarray_test.go)
    - [Maximum sum of any contiguous subarray of size k](gtci/max_subarray_test.go)
    - [Smallest subarray with a given sum](gtci/smallest_subarray_test.go)
    - [Longest substring with k distinct characters](gtci/longest_substring_k_distinct_test.go)
    - [Fruits into baskets](gtci/fruits_baskets_test.go)
    - [Longest substring without repeating characters](gtci/no_repeat_substring_test.go)
    - [Longest substring after k replacements](gtci/longest_substring_k_replacement_test.go)
    - [Longest substring after ones replacements](gtci/longest_substring_ones_replacement_test.go)
    - [Permutation in string](gtci/permutation_string_test.go)
    - [String anagrams](gtci/string_anagrams_test.go)
  - Two pointers
    - [Pair with target sum](gtci/pair_target_sum_test.go)
    - [Remove duplicates](gtci/remove_duplicates_test.go)
    - [Squaring a sorted array](gtci/square_sorted_array_test.go)
    - [Dutch national flag problem](gtci/dutch_flag_test.go)
  - Fast and slow pointers
    - [Linked list cycle](gtci/linked_list_cycle_test.go)
    - [Start of a linked list cycle](gtci/cycle_start_test.go)
    - [Happy number](gtci/happy_number_test.go)
    - [Middle of a linked list](gtci/middle_list_test.go)
    - [Palindrome linked list](gtci/palindrome_list_test.go)
    - [Reorder a linked list](gtci/reorder_list_test.go)
  - Merge intervals
    - [Merge intervals](gtci/merge_intervals_test.go)
    - [Insert interval](gtci/insert_interval_test.go)
    - [Intervals intersection](gtci/intervals_intersection_test.go)
    - [Conflicting appointment](gtci/conflict_appointment_test.go)
  - Cyclic sort
    - [Cyclic sort](gtci/cyclic_sort_test.go)
    - [Missing number](gtci/missing_number_test.go)
    - [Missing numbers](gtci/missing_numbers_test.go)
    - [Find duplicate](gtci/duplicate_test.go)
    - [Find duplicates](gtci/duplicates_test.go)
    - [Find corrupt pair](gtci/corrupt_pair_test.go)
  - In-place reversal of a linked list
    - [Reverse a linked list](gtci/reverse_list_test.go)
  - Tree breath first search
    - [Binary tree level order traversal](gtci/level_order_traversal_test.go)
    - [Reverse level order traversal](gtci/reverse_level_order_traversal_test.go)
    - [Zigzag traversal](gtci/zigzag_traversal_test.go)
    - [Level averages](gtci/level_avg_test.go)
    - [Minimum depth](gtci/min_depth_test.go)
    - [Maximum depth](gtci/max_depth_test.go)
    - [Level order successor](gtci/level_order_successor_test.go)
  - Tree depth first search
    - [Binary tree path sum](gtci/path_sum_test.go)
    - [Sum of path numbers](gtci/sum_path_test.go)
  - Two Heaps
    - [Find the median of a number stream](gtci/median_number_stream_test.go)
  - Subsets
    - [Subsets](gtci/subsets_test.go)
    - [Subsets with duplicates](gtci/subsets_duplicates_test.go)
    - [Permutations](gtci/spermutations_test.go)
- **[Grokking Dynamic Programming Patterns](https://www.educative.io/courses/grokking-dynamic-programming-patterns-for-coding-interviews)**
  - Knapsack
    - [0/1 Knapsack](gtci/knapsack_test.go)
    - [Equal subset sum partition](gtci/equal_subset_partition_test.go)
  - Fibonacci numbers
    - [Fibonacci numbers](gtci/fibonacci_numbers_test.go)
    - [Staircase](gtci/staircase_test.go)
    - [Number factors](gtci/number_factors_test.go)
- **[Practical Algorithms and Data Structures](https://bradfieldcs.com/algos/)**
- **Elements of Programming Interviews**
- **Lab (to experiment with different ideas)**
  - Data structures
    - Tree
      - [Balanced binary tree](lab/balanced_binary_tree_test.go)
      - [Binary tree traversals](lab/binary_tree_traverse_test.go)
    - Priority Queue: *TODO*
    - Bloom filter: *TODO*
    - [LRU cache](lab/lru_test.go)
  - Sorting
    - [Bubble sort](lab/bubble_sort_test.go)
    - [Selection sort](lab/selection_sort_test.go)
    - [Insertion sort](lab/insertion_sort_test.go)
    - [Merge Sort](lab/merge_sort_test.go)
    - [Quicksort](lab/quicksort_test.go)
    - [Heapsort](lab/heapsort_test.go)
    - [Counting Sort](lab/counting_sort_test.go)
    - Radix Sort: *TODO*

## Developing

Count the number of questions:
```
make count
```

Clean up, lint source files, run tests and be ready for a push:
```
make ready
```

Test only:
```
make test
```

Test and report:
```
make cover
```

Test verbose:
```
make test-verbose
```

Lint:
```
make lint
```

Clean up:
```
make clean
```

For more information:
```
make help
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/hoanhan101/algo.svg)](https://starchart.cc/hoanhan101/algo)
