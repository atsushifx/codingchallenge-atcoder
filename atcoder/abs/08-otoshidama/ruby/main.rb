# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# INPUT
N,Y = gets.chomp.split(/\s+/).map(&:to_i)

# p N, Y
# solve
isFound = false
for i in (0..N) do
  for j in (0 .. (N-i)) do
    k = N - (i + j)
    sum = i * 10000 + j * 5000 + k * 1000
    # print [i,j,k], " = ", sum, "\n"
    if (sum == Y) then
      isFound = true
      break
    end
  end
  break if isFound
end

# output
if isFound then
  print i, " ", j, " ", k, "\n"
else
  print "-1 -1 -1\n"
end
