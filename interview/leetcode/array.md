# Array 
https://leetcode.com/problems/two-sum/description/

## Answer 1
```python
#!/usr/bin/env python3

# Leetcode question: https://leetcode.com/problems/two-sum/
# Time complexity: O(n^2)

nums = [2,7,11,14]
target = 9

def two_sum(nums, target):
    for i in range(len(nums)):
        for j in range(i+1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i,j]

result = two_sum(nums, target)
print(result)
```

## Answer 2

```python
#!/usr/bin/env python3

# Leetcode question: https://leetcode.com/problems/two-sum/
# Time complexity: O(n)

nums = [2,7,11,13]
target = 15
seen = {}

def twoSum(nums, target):
    for i in range(len(nums)):
        needed = target - nums[i]
        if needed in seen:
            return [seen[needed], i]
    
        seen[nums[i]] = i

print(twoSum(nums,target))
```
