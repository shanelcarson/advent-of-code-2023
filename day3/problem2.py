def parse_and_sum(lines):
    symbol_map = dict()
    for line_idx in range(len(lines)):
        line = lines[line_idx]
        num_builder = ""
        for idx in range(len(line)):
            if line[idx].isdigit():
                num_builder += line[idx]

            if not line[idx].isdigit() or idx == len(line) - 1:
                start_idx = (
                    idx - len(num_builder) + 1
                    if line[idx].isdigit()
                    else idx - len(num_builder)
                )

                if len(num_builder) > 0:
                    add_gear_num_pair(
                        lines,
                        line_idx,
                        start_idx,
                        len(num_builder),
                        int(num_builder),
                        symbol_map,
                    )
                num_builder = ""
    return symbol_map


def add_gear_num_pair(lines, line_idx, start_idx, num_size, num_builder, symbol_map):
    for i in range(start_idx - 1, start_idx + num_size + 1):
        gear_upper = gear_near(lines, line_idx - 1, i)
        gear_lower = gear_near(lines, line_idx + 1, i)

        if gear_upper != "":
            add_to_map(symbol_map, gear_upper, num_builder)
        if gear_lower != "":
            add_to_map(symbol_map, gear_lower, num_builder)

    gear_near_left = gear_near(lines, line_idx, start_idx - 1)
    gear_near_right = gear_near(lines, line_idx, start_idx + num_size)

    if gear_near_left != "":
        add_to_map(symbol_map, gear_near_left, num_builder)
    if gear_near_right != "":
        add_to_map(symbol_map, gear_near_right, num_builder)


def gear_near(lines, line_idx, idx):
    if (
        (line_idx < 0)
        or (line_idx >= len(lines))
        or (idx < 0)
        or (idx >= len(lines[line_idx]))
    ):
        return False
    symbol = lines[line_idx][idx]
    if symbol == "*":
        return f"{line_idx},{idx}"
    return ""


def add_to_map(map, k, v):
    if k in map:
        map[k].append(v)
    else:
        map[k] = [v]


def main():
    with open("input2.txt") as f:
        lines = f.readlines()
        for line_idx in range(0, len(lines)):
            lines[line_idx] = lines[line_idx].strip()

        sum = 0
        symbol_map = parse_and_sum(lines)
        for vals in symbol_map.values():
            if len(vals) == 2:
                sum += vals[0] * vals[1]
        print(f"The sum is {sum}.")


if __name__ == "__main__":
    main()
