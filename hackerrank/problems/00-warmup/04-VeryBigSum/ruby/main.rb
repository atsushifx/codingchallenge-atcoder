#!/bin/ruby

#
# Complete the 'aVeryBigSum' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts LONG_INTEGER_ARRAY ar as parameter.
#

numCnt = gets.strip.to_i
nums = gets.strip.split(' ').map(&:to_i)

def aVeryBigSum(nums)
    sum = 0
	nums.each { |num| sum += num }
    sum
end

sum = aVeryBigSum(nums)
print sum
