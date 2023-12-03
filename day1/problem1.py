def main():
    with open("input1.txt") as f:
        sum = 0
        lines = f.readlines()
        for line in lines:
            num = ""
            for i in range(len(line)):
                if line[i].isdigit():
                    num += line[i]
                    break
            for i in reversed(range(len(line))):
                if line[i].isdigit():
                    num += line[i]
                    break
            sum += int(num)

    print("Sum is: " + str(sum))


if __name__ == "__main__":
    main()
