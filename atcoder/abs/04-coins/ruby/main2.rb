# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# main

# get Param
a = gets.to_i
b = gets.to_i
c = gets.to_i

sum = gets.to_i

count = 0
(0 ... a + 1).each { |i|
  (0 ... b + 1).each { |j|
    (0 ... c + 1).each { |k|
      x = 500*i + 100*j + 50*k

      if (x == sum) then
        count += 1
      end
      if (x > sum) then
        break
      end
    }
  }
}
p count

