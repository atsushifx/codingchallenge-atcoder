# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# INPUT
N = gets.chomp.to_i
cards = gets.chomp.split(" ").map(&:to_i)

# 降順にソート
cards.sort!.reverse!

point = 0
for i in (0 .. cards.length-1) do
  if (i%2 == 0) then
    point += cards[i]
  else
    point -= cards[i]
  end
end

p point

