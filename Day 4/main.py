# Advent of Code -- Day 4
def bingo(matrix: list) -> bool:
    def checkDown(j: int) -> bool:
        for i in range(5):
            if matrix[i][j] == False:
                return False
        return True

    def checkRight(i: int) -> bool:
        for j in range(5):
            if matrix[i][j] == False:
                return False
        return True

    def checkDownRight() -> bool:
        for i in range(5):
            if matrix[i][i] == False:
                return False
        return True

    def checkUpRight() -> bool:
        j = 0
        for i in reversed(range(5)):
            if matrix[i][j] == False:
                return False
            j += 1
        return True

    for j in range(5):
        #print(j)
        if checkDown(j):
            return True

    for i in range(5):
        if checkRight(i):
            return True
    
    if checkDownRight() or checkUpRight():
        return True

    return False


datalist = []

with open('input.in', 'r') as file:
    datalist = file.readlines()

# Turning the first line into a list of ints
drawn = [int(s) for s in datalist[0].split(',')]

# Reading in the bingo boards
boards = [] #3d
curBoard = [] #2d
for line in range(2,len(datalist)):
    if datalist[line] == '\n':
        boards.append(curBoard)
        curBoard = []
    else:
        curBoard.append([int(s) for s in datalist[line].split()])


# Creating a boolean matrix to deal with checks
checked = [ [ [False for i in range(5)] for j in range(5)] for b in range(len(boards)) ]

found = False
foundBoard = -1
foundDrawn = -1
drawnCount = 0

# These two variables are for Part 2
foundCount = 0
foundboards = [False for i in range(len(boards))]
#
for num in drawn:
    #if found:
    #    break
    drawnCount += 1

    # Going through each board and checking boxes
    for b in range(len(boards)):
        found = False
        for i in range(len(boards[b])):
            for j in range(len(boards[b][i])):
                if boards[b][i][j] == num and foundCount < 100 and foundboards[b] == False:
                    boards[b][i][j] = -1
                    checked[b][i][j] = True
                    found = bingo(checked[b])

                    if found:
                        foundBoard = b
                        foundDrawn = num

                        foundCount += 1
                        foundboards[b] = True

                if found:
                    break

            if found:
                break


# Getting sum of board
boardSum = 0
b = foundBoard
for i in range(5):
    for j in range(5):
        if boards[b][i][j] != -1:
            boardSum += boards[b][i][j]

print('BINGOOOOO')
print(boardSum)
print(foundDrawn)
print(boardSum * foundDrawn)
