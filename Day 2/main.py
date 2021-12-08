# Advent of Code -- Day 2
datalist = []

with open('input.in', 'r') as file:
    datalist = file.readlines()

# Part 1
x = y = 0
for line in range(0, len(datalist)):
    command = datalist[line].split()
    if command[0] == 'forward':
        x += int(command[1])
    elif command[0] == 'down':
        y += int(command[1])
    else:
        y -= int(command[1])

#print(x*y)


# Part 2
x = y = aim = 0
for line in range(0, len(datalist)):
    command = datalist[line].split()
    if command[0] == 'forward':
        x += int(command[1])
        y += aim * int(command[1])
    elif command[0] == 'down':
        aim += int(command[1])
    else:
        aim -= int(command[1])

print(x*y)