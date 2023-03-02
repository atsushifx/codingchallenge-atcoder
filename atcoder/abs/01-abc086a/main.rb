# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

a, b = gets.chomp.split(" ").map(&:to_i)
ans = ["Even", "Odd"]

puts ans[(a%2) * (b%2)]
