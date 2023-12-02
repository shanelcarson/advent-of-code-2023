def assessGames(lines):
    sum = 0
    for line in lines:
        cubes = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        product = 1
        smallest_cube_set(line, cubes)
        power = 1
        for val in cubes.values():
            power *= val
        sum += power
    return sum


def smallest_cube_set(line, cubes):
    game = line.split(": ")
    showings = game[1].split("; ")
    for showing in showings:
        cubes_shown = showing.split(", ")
        for cube_shown in cubes_shown:
            cube_record = cube_shown.split(" ")
            cube_color = cube_record[1].strip()
            cube_amount = int(cube_record[0])
            if cube_amount > cubes[cube_color]:
                cubes[cube_color] = cube_amount


def main():
    with open("input2.txt") as f:
        lines = f.readlines()
        print(f"Total comes to: {assessGames(lines)}")


if __name__ == "__main__":
    main()
