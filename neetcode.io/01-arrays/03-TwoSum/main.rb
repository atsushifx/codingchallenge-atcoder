#
# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
#

# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    dic = {}
    nums.each_with_index do |n, i|
        if dic[target - n]
            return [dic[target - n], i]
        end
        dic[n] = i
    end
    []
end

p two_sum([2, 7, 11, 15], 9)
p two_sum([3, 2, 4], 6)
p two_sum([3, 3], 6)
