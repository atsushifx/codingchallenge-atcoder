# Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT



## read traveling point

## travleing stop points list
class TPlist
  def initialize
    tpStart = { "t":0, "x":0, "y":0 }
    @points = [tpStart]
  end

  # getter
  def point(n)
    return @points[n]
  end

  def points
    return @points
  end

  # setter
  def add(newT, newX, newY)
    newTpoint = {"t":newT, "x":newX, "y":newY}
    @points.push(newTpoint)
  end

  # setter from input
  def readN(n)
    for i in (1 .. n) do
      (newT, newX, newY ) = gets.chomp.split(" ").map(&:to_i)
      add(newT, newX, newY)
    end
  end

  ## calc
  def distance(n)
    dt = {t:0, d:0}
    if (n >= 1) then
      p1 = @points[n]
      p0 = @points[n-1]
      dt[:t] = p1[:t] - p0[:t]
      dt[:d] = (p1[:x] - p0[:x]).abs + (p1[:y] - p0[:y]).abs
    end
    return dt
  end
end

def canMovePoints(tpList, n)
  for i in (1..n) do
    dt = tpList.distance(i)
    dt2 = dt[:t] - dt[:d]
    # p dt,dt2
    if (dt2 < 0) then
      return false
    end
    if ((dt2 % 2) > 0) then
      return false
    end
  end
  return true
end


## Main Roution
# INPUT
N = gets.to_i
tp = TPlist.new
tp.readN(N)

# solve
canMove = canMovePoints(tp, N)

# output
print (canMove ? "Yes" : "No"), "\n"

