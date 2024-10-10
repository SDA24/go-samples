Linear Search is a straightforward and basic search algorithm used to find a target value within a list. It's called "linear" because it goes through the list in a linear fashion, meaning it checks one item at a time, starting from the first element and moving towards the end. Here's an in-depth explanation:

How Linear Search Works:  
1. Sequential Search: Linear search begins at the first element of the list (or array). It compares each element with the target value you're searching for.  
2. Element-by-Element Check: If the current element matches the target value, the search ends, and the algorithm returns the index of that element. Otherwise, it moves on to the next element in the sequence.  
3. Exhaustive Search: If the algorithm reaches the end of the list without finding the target, it concludes that the target is not present in the list and returns a failure.

Performance:  
Time Complexity: O(n), where n is the number of elements in the list. The algorithm has to check each element individually, so in the worst case (when the target is at the end of the list or not present), it will need to inspect every element once.  
