import sys

def main():
    if len(sys.argv) != 2:
        print("Invalid number of command line arguments.")
        return
    file = open(sys.argv[1], 'r')
    contents = file.readlines()

    position = 0
    depth = 0

    for line in contents:
        motion, value = line.split(" ")
        value = int(value)
        if motion[0] == 'f':
            position += value
        elif motion[0] == 'd':
            depth += value
        elif motion[0] == 'u':
            depth -= value
        else:
            print("Invalid motion:", motion)

    print("Horizontal position:", position)
    print("Depth:", depth)
    print("Result:", position * depth)

if __name__ == "__main__":
    main()
