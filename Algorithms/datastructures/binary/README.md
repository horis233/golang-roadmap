# Binary

## Common binary operations

### Basic operation

a=0^a=a^0

0=a^a

Derived from the above two: a=a^b^b

### Exchange two numbers

a=a^b

b=a^b

a=a^b

### remove last 1

a=n&(n-1)

### Get last 1

diff=(n&(n-1))^n

## Common Topics

[single-number](https://leetcode-cn.com/problems/single-number/)

> Given a **non-empty** integer array, except that an element appears only once, each other element appears twice. Find the element that appears only once.

```go
func singleNumber(nums []int) int {
    // 10 ^10 == 00
    // XOR of two numbers becomes 0
    result:=0
    for i:=0;i<len(nums);i++{
        result=result^nums[i]
    }
    return result
}
```

[single-number-ii](https://leetcode-cn.com/problems/single-number-ii/)

> Given a **non-empty** integer array, except that an element appears only once, each other element appears three times. Find the element that appears only once.

```go
func singleNumber(nums []int) int {
// Count the number of 1 for each digit
var result int
for i := 0; i <64; i++ {
sum := 0
for j := 0; j <len(nums); j++ {
// count the number of 1
sum += (nums[j] >> i) & 1
}
// restore bit 00^10=10 or use | can also
result ^= (sum% 3) << i
}
return result
}
```

[single-number-iii](https://leetcode-cn.com/problems/single-number-iii/)

> Given an integer array `nums`, where exactly two elements appear only once, and all other elements appear twice. Find the two elements that appear only once.

```go
func singleNumber(nums []int) []int {
    // a=a^b
    // b=a^b
    // a=a^b
    // The key point is how to divide a^b into two parts, scheme: can be distinguished by the last 1 of diff

    diff:=0
    for i:=0;i<len(nums);i++{
        diff^=nums[i]
    }
    result:=[]int{diff,diff}
    // Remove the XOR at the end and get the position of the last 1
    diff=(diff&(diff-1))^diff
    for i:=0;i<len(nums);i++{
        if diff&nums[i]==0{
            result[0]^=nums[i]
        }else{
            result[1]^=nums[i]
        }
    }
    return result
}
```

[number-of-1-bits](https://leetcode-cn.com/problems/number-of-1-bits/)

> Write a function whose input is an unsigned integer and return the number of digits whose number is '1' in its binary expression (also known as [Haming weight](https://baike.baidu.com/item /%E6%B1%89%E6%98%8E%E9%87%8D%E9%87%8F)).

```go
func hammingWeight(num uint32) int {
    res:=0
    for num!=0{
        num=num&(num-1)
        res++
    }
    return res
}
```

[counting-bits](https://leetcode-cn.com/problems/counting-bits/)

> Given a non-negative integer **num**. For each number i in the range 0 ≤ i ≤ num, calculate the number of 1s in its binary number and return them as an array.

```go
func countBits(num int) []int {
    res:=make([]int,num+1)

    for i:=0;i<=num;i++{
        res[i]=count1(i)
    }
    return res
}
func count1(n int)(res int){
    for n!=0{
        n=n&(n-1)
        res++
    }
    return
}
```

Another dynamic programming solution

```go
func countBits(num int) []int {
    res:=make([]int,num+1)
    for i:=1;i<=num;i++{
        // The last element missing 1 can be +1
        res[i]=res[i&(i-1)]+1
    }
    return res
}
```

[reverse-bits](https://leetcode-cn.com/problems/reverse-bits/)

> Reverse the binary bits of the given 32-bit unsigned integer.

Method: turn upside down

```go
func reverseBits(num uint32) uint32 {
    var res uint32
    var pow int=31
    for num!=0{
        // Take the last digit and add it to the result after shifting to the left
        res+=(num&1)<<pow
        num>>=1
        pow--
    }
    return res
}
```

[bitwise-and-of-numbers-range](https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/)

> The given range [m, n], where 0 <= m <= n <= 2147483647, returns the bitwise AND of all numbers in this range (including m, n endpoints).

```go
func rangeBitwiseAnd(m int, n int) int {
    // m 5 1 0 1
    // 6 1 1 0
    // n 7 1 1 1
    // Turn all right shifts that may contain 0 into
    // m 5 1 0 0
    // 6 1 0 0
    // n 7 1 0 0
    // So the final result is m<<count
    var count int
    for m!=n{
        m>>=1
        n>>=1
        count++
    }
    return m<<count
}
```

## Exercise

- [ ] [single-number](https://leetcode-cn.com/problems/single-number/)
- [ ] [single-number-ii](https://leetcode-cn.com/problems/single-number-ii/)
- [ ] [single-number-iii](https://leetcode-cn.com/problems/single-number-iii/)
- [ ] [number-of-1-bits](https://leetcode-cn.com/problems/number-of-1-bits/)
- [ ] [counting-bits](https://leetcode-cn.com/problems/counting-bits/)
- [ ] [reverse-bits](https://leetcode-cn.com/problems/reverse-bits/)