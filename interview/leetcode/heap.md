# Heap
dived-and-conqure, heap, priority-queue 
https://leetcode.com/problems/k-closest-points-to-origin/description/

## Solution 1: Sort by distance
Time:  O(n log n)
Space: O(1) or O(n), depending on sorting implementation

```python
from typing import List

class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        points.sort(key=lambda p: p[0] ** 2 + p[1] ** 2)
        return points[:k]
```

## Solution 2: Heap version
Time:  O(n log n)
Space: O(n)

```python
from typing import List
import heapq

class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        heap = []

        for x, y in points:
            distance = x * x + y * y
            heapq.heappush(heap, (distance, [x, y]))

        result = []

        for _ in range(k):
            distance, point = heapq.heappop(heap)
            result.append(point)

        return result
```
