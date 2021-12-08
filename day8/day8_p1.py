import sys


def main():
    if len(sys.argv) != 2:
        print("Invalid number of command line arguments.")
        return
    file = open(sys.argv[1], 'r')
    contents = file.readlines()

    count = 0

    for line in contents:
        pattern, output = line.split("|")
        digits = output.split()
        for digit in digits:
            if len(digit) in [2, 3, 4, 7]:
               count += 1
            
        
    print("Total:", count)
  
if __name__ == "__main__":
    main()
