# Dynamic Programing

## Longest Palindromic Substring
dynamic-programming
https://leetcode.com/problems/longest-palindromic-substring/description/

### Solution 1
Time:  O(n²)
Space: O(1)

```python
class Solution:
    def longestPalindrome(self, s: str) -> str:
        start = 0
        end = 0

        def expand(left, right):
            while left >= 0 and right < len(s) and s[left] == s[right]:
                left -= 1
                right += 1

            # after loop, left and right are one step too far
            return left + 1, right - 1

        for i in range(len(s)):
            # Odd length palindrome
            left1, right1 = expand(i, i)

            # Even length palindrome
            left2, right2 = expand(i, i + 1)

            if right1 - left1 > end - start:
                start, end = left1, right1

            if right2 - left2 > end - start:
                start, end = left2, right2

        return s[start:end + 1]
```

```go
func longestPalindrome(s string) string {
	start := 0
	end := 0

	expand := func(left int, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		// left and right moved one step too far
		return left + 1, right - 1
	}

	for i := 0; i < len(s); i++ {
		// Odd length palindrome
		left1, right1 := expand(i, i)

		// Even length palindrome
		left2, right2 := expand(i, i+1)

		if right1-left1 > end-start {
			start = left1
			end = right1
		}

		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}

	return s[start : end+1]
}
```

## Generate Prentheses
dynamic-programming
https://leetcode.com/problems/generate-parentheses/description

### Solution 1
Time:  O(4^n / sqrt(n))
Space: O(n)

```python
from typing import List

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        result = []

        def backtrack(current, open_count, close_count):
            if len(current) == 2 * n:
                result.append(current)
                return

            if open_count < n:
                backtrack(current + "(", open_count + 1, close_count)

            if close_count < open_count:
                backtrack(current + ")", open_count, close_count + 1)

        backtrack("", 0, 0)
        return result
```

```go
func generateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(current string, openCount int, closeCount int)

	backtrack = func(current string, openCount int, closeCount int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if openCount < n {
			backtrack(current+"(", openCount+1, closeCount)
		}

		if closeCount < openCount {
			backtrack(current+")", openCount, closeCount+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
```

## Jump Game II
dynamic-programming, greedy
https://leetcode.com/problems/jump-game-ii/description/

### Solution 1
Time:  O(n)
Space: O(1)

```python
from typing import List

class Solution:
    def jump(self, nums: List[int]) -> int:
        jumps = 0
        current_end = 0
        farthest = 0

        for i in range(len(nums) - 1):
            farthest = max(farthest, i + nums[i])

            if i == current_end:
                jumps += 1
                current_end = farthest

        return jumps
```

```go
func jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}
```
