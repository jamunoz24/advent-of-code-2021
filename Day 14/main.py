# Advent of code -- Day 14
import sys

sys.stdin = open('input.in')
data = sys.stdin.read().splitlines()

# Reading in our data
template = data[0]
instructions = {}
for i in range(1, len(data)):
    if data[i]:
        pair = data[i].split(' -> ')
        instructions[pair[0]] = pair[1]

steps = 10

# Part 1:
# Doing the shuffle
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
        tempCounts[char] = 1

maxChar = 0
minChar = pow(2,32)-1
for key in tempCounts:
    val = tempCounts[key]
    maxChar = max(maxChar, val)
    minChar = min(minChar, val)

difference = maxChar - minChar
print(tempCounts)
print("After", steps,"steps:", difference)

steps = 40

# Part 2:
# Optimizing our shuffle
template = data[0]
pairCounts = {}

# First storing the template in the dictionary
for i in range (1, len(template)):
    pair = template[i-1] + template[i]
    if pair in pairCounts:
        pairCounts[pair] += 1
    else:
        pairCounts[pair] = 1


# Now, doing the shuffle with the dictionary
tempCounts = {}
for char in template:
    if char in tempCounts:
        tempCounts[char] += 1
    else:
        tempCounts[char] = 1

for _ in range(steps):
    newDict = pairCounts.copy()
    for pair in pairCounts:
        if pairCounts[pair] > 0 and pair in instructions:
            # Adding the pair to the new dict
            occurrences = pairCounts[pair]
            newDict[pair] -= occurrences
            newChar = instructions[pair]
            newPair1 = pair[0] + newChar
            newPair2 = newChar + pair[1]
            if newPair1 in newDict:
                newDict[newPair1] += occurrences
            else:
                newDict[newPair1] = occurrences
            if newPair2 in newDict:
                newDict[newPair2] += occurrences
            else:
                newDict[newPair2] = occurrences

            # Adding to the character count
            if newChar in tempCounts:
                tempCounts[newChar] += occurrences
            else:
                tempCounts[newChar] = occurrences
                
    pairCounts = newDict.copy()


maxChar = tempCounts['N']
minChar = tempCounts['N']
for key in tempCounts:
    val = tempCounts[key]
    maxChar = max(maxChar, val)
    minChar = min(minChar, val)

difference = maxChar - minChar
print()
print(tempCounts)
print("After", steps,"steps:", difference)
