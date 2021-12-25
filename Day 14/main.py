# Advent of code -- Day 14
import sys

sys.stdin = open('example.in')
data = sys.stdin.read().splitlines()

# Reading in our data
template = data[0]
instructions = {}
for i in range(1, len(data)):
    if data[i]:
        pair = data[i].split(' -> ')
        instructions[pair[0]] = pair[1]


# Part 1:
# Doing the shuffle
steps = 10
for _ in range(steps):
    newTemp = ''
    for i in range(1, len(template)):
        pair = template[i-1] + template[i]
        if pair in instructions:
            newTemp += template[i-1] + instructions[pair]
    
    newTemp += template[-1]
    template = newTemp

# Finding the most and least occuring character
tempCounts = {}
for char in template:
    if char in tempCounts:
        tempCounts[char] += 1
    else:
        tempCounts[char] = 0

maxChar = 0
minChar = pow(2,32)-1
for key in tempCounts:
    val = tempCounts[key]
    maxChar = max(maxChar, val)
    minChar = min(minChar, val)

difference = maxChar - minChar
print("After 10 Steps:", difference)


# Part 2:
# Optimizing our shuffle

