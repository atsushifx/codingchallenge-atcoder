# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# INPUT
N = gets.chomp.to_i
mochi = []
for i in (0 ... N) do
  mochi[i] = gets.chomp.to_i
end

# solve
kagamimochi = mochi.sort.uniq

# output
p kagamimochi.size
