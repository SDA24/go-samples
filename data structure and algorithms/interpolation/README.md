Interpolation Search Algorithm:  
    Interpolation Search is an improvement over Binary Search when the elements in a sorted array are uniformly distributed. While Binary Search always   checks the middle element, Interpolation Search estimates the position of the target element by considering the distribution of elements and    "interpolating" to find a closer guess.

Key Concept:  
The formula used to estimate the position is:  
        $$
            pos = low + ((x - arr[low]) * (high - low))/(arr[high] - arr[low])
        $$