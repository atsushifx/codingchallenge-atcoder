# Copyright (c) 2024 Furukawa, Atsushi <atsushifx@gmail.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

def odd_or_even(number)

    # add your code here
	(number %2) == 0
end

## check number :: main
cnt = gets.to_i

for i in 1 .. cnt
	num = gets.to_i
	print odd_or_even(num), "\n"
end
