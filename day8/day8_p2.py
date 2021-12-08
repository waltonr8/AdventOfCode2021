import sys

def isNumInStr(number, string):
    for char in number:
        if char not in string:
            return False
    return True

def getBottomLeft(eight, nine):
    for char in eight:
        if char not in nine:
            return char
    
    # should never get here
    return ''

def classifyLit6(numValues, lit6):
    for val in lit6:
        if isNumInStr(numValues[7], val) and isNumInStr(numValues[4], val):
            numValues[9] = val
        elif isNumInStr(numValues[7], val):
            numValues[0] = val
        else:
            numValues[6] = val

def classifyLit5(numValues, lit5):
    bottomLeft = getBottomLeft(numValues[8], numValues[9])
    for val in lit5:
        if isNumInStr(numValues[7], val):
            numValues[3] = val
        elif bottomLeft in val:
            numValues[2] = val
        else:
            numValues[5] = val

def getNumberVals(patterns):
    numValues = ['', '', '', '', '', '', '', '', '', '']
    lit5 = []
    lit6 = []
    patterns = patterns.split()
    for digit in patterns:
        if len(digit) == 2:
            numValues[1] = "".join(sorted(digit))
        elif len(digit) == 3:
            numValues[7] = "".join(sorted(digit))
        elif len(digit) == 4:
            numValues[4] = "".join(sorted(digit))
        elif len(digit) == 5:
            lit5.append("".join(sorted(digit)))
        elif len(digit) == 6:
            lit6.append("".join(sorted(digit)))
        elif len(digit) == 7:
            numValues[8] =  "".join(sorted(digit))

    classifyLit6(numValues, lit6)
    classifyLit5(numValues, lit5)    
    return numValues

def main():
    if len(sys.argv) != 2:
        print("Invalid number of command line arguments.")
        return
    file = open(sys.argv[1], 'r')
    contents = file.readlines()
    count = 0

    for line in contents:
        pattern, output = line.split("|")
        
        numVals = getNumberVals(pattern)
        digits = output.split()
        number = []
        for digit in digits:
            digit = "".join(sorted(digit))
            for i in range(len(numVals)):
                if numVals[i] == digit:
                    number.append(i)

        for i in range(len(number), 0, -1):
            count += number[len(number) - i] * (10 ** (i - 1))
            
        
    print("Total:", count)
  
if __name__ == "__main__":
    main()
