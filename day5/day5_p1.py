import sys

def getPoints(line):
    start, end = line.split("->")
    startX, startY = start.strip().split(",")
    endX, endY = end.strip().split(",")
    return int(startX), int(startY), int(endX), int(endY)

def growMap(m, size, startX, startY, endX, endY):
    # need more rows
    if startY > size[0]:
        for i in range(startY - size[0]):
            m.append([0 for i in range(size[1] + 1)])
        size[0] = startY
    if endY > size[0]:
        for i in range(endY - size[0]):
            m.append([0 for i in range(size[1] + 1)])
        size[0] = endY

    # need more cols
    if startX > size[1]:
        for i in range(size[0] + 1):
            m[i].extend([0 for j in range(startX - size[1])])
        size[1] = startX
    if endX > size[1]:
        for i in range(size[0] + 1):
            m[i].extend([0 for j in range(endX - size[1])])
        size[1] = endX

def drawLines(m, startX, startY, endX, endY):
    if startX == endX:
        if startY > endY:
            startY, endY = endY, startY
        for i in range(startY, endY + 1):
            m[i][startX] += 1
    elif startY == endY:
        if startX > endX:
            startX, endX = endX, startX
        for i in range(startX, endX + 1):
            m[startY][i] += 1

def getCount(m, size):
    count = 0
    for i in range(size[0] + 1):
        for j in range(size[1] + 1):
            if m[i][j] >= 2:
                count += 1
    return count

def main():
    if len(sys.argv) != 2:
        print("Invalid number of command line arguments.")
        return
    file = open(sys.argv[1], 'r')
    contents = file.readlines()

    m = [[0]]
    size = [0, 0]

    for line in contents:
        start, end = line.split("->")
        startX, startY = start.strip().split(",")
        endX, endY = end.strip().split(",")

        startX, startY, endX, endY = getPoints(line)
        growMap(m, size, startX, startY, endX, endY)
        drawLines(m, startX, startY, endX, endY)
        
    count = getCount(m, size)
    print("Total:", count)
  
if __name__ == "__main__":
    main()
