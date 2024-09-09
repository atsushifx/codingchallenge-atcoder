#!/usr/bin/env ruby

# read dialog
cnt = gets.strip.to_i
dlgNums = Array.new(cnt) { Array.new(cnt, 0) }
for i in 0 .. cnt - 1 do
  nums = gets.strip.split.map(&:to_i)
	for j in 0 .. cnt -1 do
	    dlgNums[i][j] = nums[j]
	end
end

# primary
p_sum = 0
for i in 0 .. cnt - 1 do
  p_sum += dlgNums[i][i]
end

# secondary
s_sum = 0
for i in 0 .. cnt - 1 do
	j = (cnt -1) -i
  s_sum += dlgNums[i][j]
end

print (p_sum - s_sum).abs
