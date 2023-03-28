# @param {String[]} strs
# @return {String[][]}
def group_anagrams(strs)
  list1 = {}
  strs.each do |str|
    h = shash(str)
    list1[h] = [] unless list1[h]
    list1[h].push(str)
  end
  list1.values
end

def shash(str)
  str.bytes.sort.hash
end
