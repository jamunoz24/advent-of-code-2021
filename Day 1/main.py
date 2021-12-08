# Advent of Code -- Day 1
datalist = []

with open('input.in', 'r') as file:
    datalist = file.readlines()

for line in range(0, len(datalist)):
    datalist[line] = int(datalist[line])


# Part 1
count = 0
for i in range(1, len(datalist)):
    if datalist[i] > datalist[i-1]:
        count += 1



# Part 2
count = 0
for i in range(0, len(datalist)-3):
    midSums = datalist[i+1] + datalist[i+2]
    firstSum = datalist[i] + midSums
    secondSum = midSums + datalist[i+3]

    if firstSum < secondSum:
        count += 1

print(count)