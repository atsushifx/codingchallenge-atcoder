#
# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
#

# @param {String} s
# @param {String} t
# @return {Boolean}
def is_anagram(src, target)
  h1 = src.unpack('U*').sort
  h2 = target.unpack('U*').sort
  h1 == h2
end
