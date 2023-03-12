# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# abc_08b : coins

# COIN定数
C500 = 500
C100 = 100
C50  = 50


# Methods
# get 1st coin pay
def getFirstPays(coins, sum)
  pays = {}
  coins.each { |coin, count|
    c = sum / coin
    if (c > count) then
      c = count
    end
    pays[coin] = c
    sum -= coin * c
  }
  if (sum != 0) then
    return nil
  end
  return pays
end

# 100 => 50 * 2
def pay100to50(coins, pays)
  count = 0
  while pays[C100] > 0 do

    if (coins[C50] < 2) then
      break
    end
    pays[C100]  -= 1
    pays[C50]   += 2
    coins[C50]  -= 2
    coins[C100] += 1
    count += 1

    p  pays
  end
  return count
end

# 500 => 100 * 5
def pay500to100(coins, pays)
  count = 0
  while pays[C500] > 0
    count += pay100to50(coins, pays)

    if (coins[C100] < 5) then
      break
    end
    pays[C500]  -= 1
    pays[C100]  += 5
    coins[C500] += 1
    coins[C100] -= 5
    count += 1
  end

  count += pay100to50(coins, pays)
  return count
end


# main

# get Param
coins = {}
coins[C500] = gets.to_i
coins[C100] = gets.to_i
coins[C50]  = gets.to_i

sum = gets.to_i


# solve
pays = getFirstPays(coins, sum)
p pays, "---"

count = 0
if (pays == nil) then # no counts
  count = 0
else
  count = 1
  p pays
  count += pay500to100(coins, pays)
end


p count
