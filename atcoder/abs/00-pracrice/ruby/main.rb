#!/usr/bin/env ruby
#
# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
#

a = gets.to_i
b, c = gets.chomp.split(" ").map(&:to_i)
s = gets.chomp

sum = a+b+c

puts "#{sum} #{s}"
