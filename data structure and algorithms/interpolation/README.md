Interpolation Search Algorithm:  
    Interpolation Search is an improvement over Binary Search when the elements in a sorted array are uniformly distributed. While Binary Search always   checks the middle element, Interpolation Search estimates the position of the target element by considering the distribution of elements and    "interpolating" to find a closer guess.

Key Concept:  
The formula used to estimate the position is:  
            pos = low + ((x - arr[low]) * (high - low))/(arr[high] - arr[low])

Here’s how each term contributes to the formula:  
1. low: This is the starting index of the search range in the array.  
2. high: This is the ending index of the search range in the array.  
3. arr[low] and arr[high]: These are the values of the array elements at the low and high indices. These values help determine the range of possible values for the search.  
4. x: This is the target value we are searching for.