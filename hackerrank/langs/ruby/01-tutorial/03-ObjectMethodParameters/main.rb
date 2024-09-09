# Copyright (c) 2024 Furukawa, Atsushi <atsushifx@gmail.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
## test main
cnt = gets.to_i
for i in 1 .. cnt
	a, b, c = gets.chomp.split.map(&:to_i)
	print a.between?(b, c), "\n"
end
