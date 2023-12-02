import re


def main():
    sum = 0
    with open("input2.txt") as f:
        lines = f.readlines()
        num = []
        for l in lines:
            rep = {
                "one": "one1one",
                "two": "two2two",
                "three": "three3three",
                "four": "four4four",
                "five": "five5five",
                "six": "six6six",
                "seven": "seven7seven",
                "eight": "eight8eight",
                "nine": "nine9nine",
            }

            rep = dict((re.escape(k), v) for k, v in rep.items())

            line = l
            for key in rep.keys():
                line = re.sub(key, rep[key], line)

            num = []
            for i in range(len(line)):
                if line[i].isdigit():
                    num.append(line[i])

            line_sum = int(num[0] + num[-1])
            sum += line_sum

    print("Sum is: " + str(sum))


if __name__ == "__main__":
    main()
