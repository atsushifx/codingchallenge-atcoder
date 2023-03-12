# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT



## functions
def addNums(n)
  nList = n.to_s.chars.map(&:to_i)
  return nList.sum
end

# main

# INPUT
N, A, B = gets.split(/\s+/).map(&:to_i)

# solve
total = 0
for i in (1... N+1) do
  x = addNums(i)
  if ((A<=x)&&(x<=B)) then
    total += i
    #   p "#{i}=>#{x}"
  end
end

# OUTPUT
p total
