def assessGames(lines, cubes):
    sum = 0
    for line in lines:
        sum += checkGame(line, cubes)
    return sum


def checkGame(line, cubes):
    split_lines = line.split(": ")
    id = int(split_lines[0][5:])
    showings = split_lines[1].split("; ")
    for showing in showings:
        cubes_shown = showing.split(", ")
        for cube_shown in cubes_shown:
            cube_record = cube_shown.split(" ")
            if int(cube_record[0]) > cubes[cube_record[1].strip()]:
                return 0
    return id


def main():
    cubes = {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    with open("input1.txt") as f:
        lines = f.readlines()
        print(f"Total comes to: {assessGames(lines, cubes)}")


if __name__ == "__main__":
    main()
