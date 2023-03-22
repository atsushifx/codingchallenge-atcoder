# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

## initialize
Keywords = [
  "dream",
  "dreamer",
  "erase",
  "eraser",
]

## solve


## solver function
def isMatchKeywords(str)
  ret = false
  if (str == "") then
    ret = true
  else
    ret = false
    Keywords.each do |k|
      idx = str.index(k)
      if (idx === 0) then
        ret2 = isMatchKeywords(str.delete_prefix(k))
        if (ret2) then
          ret = true
          break
        else

        end
      end
    end
  end
  return ret
end

## main
s = gets.chomp
print (isMatchKeywords(s) ? "YES" : "NO"),"\n"

