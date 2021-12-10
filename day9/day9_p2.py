import sys
from functools import reduce

def printHeatmap(heatMap):
    for i in range(len(heatMap)):
        for j in range(len(heatMap[i])):
            print(heatMap[i][j], end=" ")
        print()

def getLowPoints(heatMap):
    lowPoints = []
    for i in range(len(heatMap)):
        for j in range(len(heatMap[i])):
            val = heatMap[i][j]
            up = i - 1
            down = i + 1
            left = j - 1
            right = j + 1
            if up >= 0 and val >= heatMap[up][j]:
                continue
            if down < len(heatMap) and val >= heatMap[down][j]:
                continue
            if left >= 0 and val >= heatMap[i][left]:
                continue
            if right < len(heatMap[i]) and val >= heatMap[i][right]:
                continue
            
            lowPoints.append((val, i, j))

    return lowPoints

def check(previousVal, i, j, heatMap):
    if i < 0 or i >= len(heatMap) or j < 0 or j >= len(heatMap[0]):
        # invalid index
        return 0 
    
    if previousVal < heatMap[i][j] and heatMap[i][j] != -1 and heatMap[i][j] != 9:
        up = check(heatMap[i][j], i - 1, j, heatMap)
        down = check(heatMap[i][j], i + 1, j, heatMap)
        left = check(heatMap[i][j], i, j - 1, heatMap)
        right = check(heatMap[i][j], i, j + 1, heatMap)
        heatMap[i][j] = -1
        return up + down + left + right + 1
    
    return 0
        
def getBasins(lowPoints, heatMap):
    basins = []
    for val in lowPoints:
        point, startX, startY = val
        up = check(heatMap[startX][startY], startX - 1, startY, heatMap)
        down = check(heatMap[startX][startY], startX + 1, startY, heatMap)
        left = check(heatMap[startX][startY], startX, startY - 1, heatMap)
        right = check(heatMap[startX][startY], startX, startY + 1, heatMap)
        basins.append(up + down + left + right + 1)
        heatMap[startX][startY] = -1

    return basins

def main():
    if len(sys.argv) != 2:
        print("Invalid number of command line arguments.")
        return
    file = open(sys.argv[1], 'r')
    contents = file.readlines()

    count = 0

    heatMap = []
    for line in contents:
        heatMap.append([])
        for char in line:
            if char != '\n':
                heatMap[-1].append(int(char))

    lowPoints = getLowPoints(heatMap)
    
    basins = getBasins(lowPoints, heatMap)

    basins = sorted(basins, reverse=True)[:3]

    print("Total:", reduce((lambda x, y: x * y), basins))
  
if __name__ == "__main__":
    main()
