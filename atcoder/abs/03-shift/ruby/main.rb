# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# get param
n = gets.chomp.to_i
nums = gets.chomp.split(" ").map(&:to_i)



# shift number
isEven = true
count = 0
while isEven do
  nums = nums.map { |num|
    if ((num % 2) == 1)
      isEven = false
      break
    end
    num / 2
  }
  if (isEven)
    count += 1
  end
end

p count



