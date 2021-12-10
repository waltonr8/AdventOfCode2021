import sys

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
            
            lowPoints.append(val)

    return lowPoints
            

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
    print(lowPoints)
            
        
    print("Total:", sum(lowPoints) + len(lowPoints))
  
if __name__ == "__main__":
    main()
