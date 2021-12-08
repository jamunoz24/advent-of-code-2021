# Advent of Code -- Day 3
datalist = []

with open('input.in', 'r') as file:
    datalist = file.readlines()

# The number of bits we're working with
bits = len(datalist[0])-1

# Keeping track of the most popular bits
bitCounts = [ [0] * bits for i in range(2) ]
for line in range(0, len(datalist)):
    for i in range(0, bits):
        bitCounts[int(datalist[line][i])][i] += 1


# Calculating gamma and epsilon
gamma = ''
epsilon = ''
for i in range(bits):
    if bitCounts[0][i] > bitCounts[1][i]:
        gamma += '0'
        epsilon += '1'
    else:
        gamma += '1'
        epsilon += '0'

# binary to decimal
gamma = int(gamma,2)
epsilon = int(epsilon,2)
print('Part 1:')
print(gamma * epsilon)


# Part 2
# Keeping a list of the lines with the most popular bits
mostPop = []

# Getting our starting list
firstBit = ''
if bitCounts[0][0] > bitCounts[1][0]:
    firstBit = '0'
else:
    firstBit = '1'

for line in range(len(datalist)):
    if datalist[line][0] == firstBit:
        mostPop.append(datalist[line].strip())

# Getting the most popular bit then filtering the list with it
for i in range(1, bits):
    if len(mostPop) == 1:
        break

    mostCommon = [0,0]
    for line in range(len(mostPop)):
        if mostPop[line][i] == '0':
            mostCommon[0] += 1
        else:
            mostCommon[1] += 1
    
    if mostCommon[0] > mostCommon[1]:
        mostCommon = '0'
    else:
        mostCommon = '1'

    newList = mostPop.copy()
    counter = 0
    for line in range(len(mostPop)):
        if mostPop[line][i] != mostCommon:
            newList[line] = ''

    while ('' in newList):
        newList.remove('')
            
    mostPop = newList.copy()


# Keeping a list of the lines with the least popular bits
leastPop = []

# getting first list
firstBit = ''
if bitCounts[0][0] <= bitCounts[1][0]:
    firstBit = '0'
else:
    firstBit = '1'

for line in range(len(datalist)):
    if datalist[line][0] == firstBit:
        leastPop.append(datalist[line].strip())

# Getting the least popular bit then filtering the list with it
for i in range(1, bits):
    if len(leastPop) == 1:
        break

    mostCommon = [0,0]
    for line in range(len(leastPop)):
        if leastPop[line][i] == '0':
            mostCommon[0] += 1
        else:
            mostCommon[1] += 1
    
    leastCommon = ''
    if mostCommon[0] <= mostCommon[1]:
        leastCommon = '0'
    else:
        leastCommon = '1'

    newList = leastPop.copy()
    counter = 0
    for line in range(len(leastPop)):
        if leastPop[line][i] != leastCommon:
            newList[line] = ''

    while ('' in newList):
        newList.remove('')
            
    leastPop = newList.copy()
    #print(mostPop) 

#print(mostPop)
#print(leastPop)
print('\nPart 2:')
print(int(mostPop[0],2) * int(leastPop[0],2))