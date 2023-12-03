def parse_and_sum(lines):
    sum = 0
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

                if (len(num_builder) > 0) and symbol_around_num(
                    lines, line_idx, start_idx, len(num_builder)
                ):
                    sum += int(num_builder)
                num_builder = ""
    return sum


def symbol_around_num(lines, line_idx, start_idx, num_size):
    for i in range(start_idx - 1, start_idx + num_size + 1):
        if symbol_near(lines, line_idx - 1, i) or symbol_near(lines, line_idx + 1, i):
            return True

    if symbol_near(lines, line_idx, start_idx - 1) or symbol_near(
        lines, line_idx, start_idx + num_size
    ):
        return True
    return False


def symbol_near(lines, line_idx, idx):
    if (
        (line_idx < 0)
        or (line_idx >= len(lines))
        or (idx < 0)
        or (idx >= len(lines[line_idx]))
    ):
        return False
    symbol = lines[line_idx][idx]
    return (not symbol.isdigit()) and (not symbol == ".")


def main():
    with open("input1.txt") as f:
        lines = f.readlines()
        for line_idx in range(0, len(lines)):
            lines[line_idx] = lines[line_idx].strip()

        sum = parse_and_sum(lines)
        print(f"The sum is {sum}.")


if __name__ == "__main__":
    main()
