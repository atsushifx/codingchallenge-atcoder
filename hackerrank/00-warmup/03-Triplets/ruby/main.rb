#!/bin/ruby
#
# Complete the 'compareTriplets' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER_ARRAY b
#


alice = gets.strip.split(' ').map(&:to_i)
bob = gets.strip.split(' ').map(&:to_i)

a_pt = 0
b_pt = 0
for i in 0 .. alice.size-1 do
	if alice[i] > bob[i] 
	    a_pt += 1
	elsif alice[i] < bob[i]
		b_pt += 1
	end
end

print a_pt, ' ', b_pt
