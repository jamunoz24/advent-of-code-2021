# Advent of Code -- Day 4
import re
datalist = []

with open('input.in', 'r') as file:
    datalist = file.readlines()


lineCoords = []
for line in datalist:
    # This could probably be all done in one line
    coordList = re.split(' -> |\n', line)
    startCoord = [int(x) for x in coordList[0].split(',')]
    endCoord = [int(x) for x in coordList[1].split(',')]

    lineCoords.append([tuple(startCoord), tuple(endCoord)])

print('line:',lineCoords[0])
print('coord:',lineCoords[0][0])

# Getting the largest x and y
largestX = 0
largestY = 0
for line in lineCoords:
    for coord in line:
        largestX = max(largestX, coord[0])
        largestY = max(largestY, coord[1])

print('largest x:', largestX)
print('largest y:', largestY)

# Creating our map
ventMap = [ [0 for _ in range(largestX+1)] for _ in range(largestY+1) ]

# Drawing the lines on the map
for line in lineCoords:
    x1, y1 = line[0][0], line[0][1]
    x2, y2 = line[1][0], line[1][1]
    # Vertical line
    if x1 == x2:
        if y1 > y2:
            for i in range(y2,y1+1):
                ventMap[x1][i] += 1
        else:
            for i in range(y1,y2+1):
                ventMap[x1][i] += 1

    # Horizontal line
    elif y1 == y2:
        if x1 > x2:
            for i in range(x2,x1+1):
                ventMap[i][y1] += 1
        else:
            for i in range(x1, x2+1):
                ventMap[i][y1] += 1

    # Part 2:
    # Diagonal lines
    else:
        # Swapping points if first point is on the right side
        if x1 > x2:
            x1, y1, x2, y2 = x2, y2, x1, y1

        difference = x2-x1
        # Going from bottom-left up
        if y1 > y2:
            for i in range(difference+1):
                ventMap[x1+i][y1-i] += 1

        # Going from top-left down
        else:
            for i in range(difference+1):
                ventMap[x1+i][y1+i] += 1


overlaps = 0
for row in ventMap:
    for num in row:
        if num > 1:
            overlaps += 1

print(overlaps)