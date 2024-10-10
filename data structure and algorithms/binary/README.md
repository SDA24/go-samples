Binary Search is a much more efficient search algorithm compared to Linear Search, especially for large datasets. It works by repeatedly dividing the search range in half. However, it requires that the list (or array) being searched is already sorted.

How Binary Search Works:  
1. Middle Element Comparison: Binary Search starts by examining the middle element of the sorted list.  
    * If the middle element matches the target, the search is complete.  
    * If the middle element is greater than the target, the search continues on the left half of the list (i.e., the smaller elements).  
    * If the middle element is smaller than the target, the search continues on the right half of the list (i.e., the larger elements).  
2. Repeat: This process of halving the search space continues until either the target element is found or the search space is exhausted.  

Performance:  
    * Time Complexity: O(log n), where n is the number of elements. This is because the search space is halved with each iteration, making it very efficient for large datasets.  
    * Space Complexity: O(1) for the iterative approach, as it requires only a few extra variables. A recursive implementation might take up additional space due to the call stack, resulting in O(log n) space complexity.  
