# Sliding window

## Template

```cpp
/* Sliding window algorithm framework */
void slidingWindow(string s, string t) {
    unordered_map<char, int> need, window;
    for (char c: t) need[c]++;

    int left = 0, right = 0;
    int valid = 0;
    while (right <s.size()) {
        // c is the character that will be moved into the window
        char c = s[right];
        // move window right
        right++;
        // Perform a series of updates of data in the window
        ...

        /*** Location of debug output ***/
        printf("window: [%d, %d)\n", left, right);
        /********************/

        // Determine whether to shrink the left window
        while (window needs shrink) {
            // d is the character that will move out of the window
            char d = s[left];
            // move window left
            left++;
            // Perform a series of updates of data in the window
            ...
        }
    }
}
```

Need to change

- 1. The window data is updated after the right pointer moves to the right
- 2. Determine whether the window should shrink
- 3. The window data is updated after the left pointer moves to the right
- 4. Calculate the result according to the meaning of the question

## Example

[minimum-window-substring](https://leetcode-cn.com/problems/minimum-window-substring/)

> Give you a string S, a string T, please find in the string S: the smallest substring containing all the letters of T

```go
func minWindow(s string, t string) string {
    // Save the sliding window character set
    win := make(map[byte]int)
    // Save the required character set
    need := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        need[t[i]]++
    }
    // window
    left := 0
    right := 0
    // match times
    match := 0
    start := 0
    end := 0
    min := math.MaxInt64
    var c byte
    for right < len(s) {
        c = s[right]
        right++
        // In the required character set, add it to the window character set
        if need[c] != 0 {
            win[c]++
            // If the current number of characters matches the required number of characters, the match value +1
            if win[c] == need[c] {
                match++
            }
        }

        // When all the characters match, start to shrink the window
        for match == len(need) {
            // Get results
            if right-left < min {
                min = right - left
                start = left
                end = right
            }
            c = s[left]
            left++
            // The left pointer points to the character set that is not needed and skips directly
            if need[c] != 0 {
                if win[c] == need[c] {
                    match--
                }
                win[c]--
            }
        }
    }
    if min == math.MaxInt64 {
        return ""
    }
    return s[start:end]
}
```

[permutation-in-string](https://leetcode-cn.com/problems/permutation-in-string/)

> Given two strings **s1** and **s2**, write a function to determine whether **s2** contains the arrangement of **s1 **.

```go
func checkInclusion(s1 string, s2 string) bool {
    win := make(map[byte]int)
    need := make(map[byte]int)
    for i := 0; i < len(s1); i++ {
        need[s1[i]]++
    }
    left := 0
    right := 0
    match := 0
    for right < len(s2) {
        c := s2[right]
        right++
        if need[c] != 0 {
            win[c]++
            if win[c] == need[c] {
                match++
            }
        }
        // When the window length is greater than the string length, shrink the window
        for right-left >= len(s1) {
            // When the window length matches the string, and the number of characters in it also matches, the condition is met
            if match == len(need) {
                return true
            }
            d := s2[left]
            left++
            if need[d] != 0 {
                if win[d] == need[d] {
                    match--
                }
                win[d]--
            }
        }
    }
    return false
}

```

[find-all-anagrams-in-a-string](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

> Given a string **s ** and a non-empty string **p**, find all substrings of **p ** in **s **, and return these substrings The starting index of.

```go
func findAnagrams(s string, p string) []int {
    win := make(map[byte]int)
need := make(map[byte]int)
for i := 0; i <len(p); i++ {
need[p[i]]++
}
left := 0
right := 0
match := 0
    ans:=make([]int,0)
for right <len(s) {
c := s[right]
right++
if need[c] != 0 {
win[c]++
if win[c] == need[c] {
match++
}
}
// When the window length is greater than the string length, shrink the window
for right-left >= len(p) {
// When the window length matches the string, and the number of characters in it also matches, the condition is met
if right-left == len(p)&& match == len(need) {
ans=append(ans,left)
}
d := s[left]
left++
if need[d] != 0 {
if win[d] == need[d] {
match--
}
win[d]--
}
}
}
return ans
}
```

[longest-substring-without-repeating-characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

> Given a string, please find out the length of the longest substring that does not contain repeated characters.
> Example 1:
>
> Input: "abcabcbb"
> Output: 3
> Explanation: Because the longest substring without repeated characters is "abc", its length is 3.

```go
func lengthOfLongestSubstring(s string) int {
    // The core points of the sliding window: 1. The right pointer moves to the right 2. The window is contracted according to the meaning of the question 3. The left pointer moves to the right to update the window 4. The result is calculated according to the meaning of the question
    if len(s)==0{
        return 0
    }
    win:=make(map[byte]int)
    left:=0
    right:=0
    ans:=1
    for right<len(s){
        c:=s[right]
        right++
        win[c]++
        // Shrink the window
        for win[c]>1{
            d:=s[left]
            left++
            win[d]--
        }
        // Calculation results
        ans=max(right-left,ans)
    }
    return ans
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

## to sum up

- Similar to the dual pointer problem, it is more like an upgraded version of the dual pointer. The core point of the sliding window is to maintain a window set and process it according to the window set
- Core steps
  - right
  - Shrink
  - left move right
  - Seek results

## Exercise

- [ ] [minimum-window-substring](https://leetcode-cn.com/problems/minimum-window-substring/)
- [ ] [permutation-in-string](https://leetcode-cn.com/problems/permutation-in-string/)
- [ ] [find-all-anagrams-in-a-string](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)
- [ ] [longest-substring-without-repeating-characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)