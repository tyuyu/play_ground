+解题思路
0~i 的和记为 f(i) 
可知  f(i) = f(i-1) + nums[i]
若 j~i 的和 为k 
有 
f(j-1) + k = f(i) 
f(j-1) = f(i) - k 
统计 f(j-1) 和为 f(i) - k 的个数
等价于 统计 f(i) - k 个数
