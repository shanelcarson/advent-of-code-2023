from queue import Queue

new_dir = {
    "N": {
        "7": "W",
        "F": "E",
    },
    "S": {
        "L": "E",
        "J": "W",
    },
    "E": {
        "7": "S",
        "J": "N",
    },
    "W": {
        "L": "N",
        "F": "S",
    },
}

new_dir_inc = {
    "N": [-1, 0],
    "S": [1, 0],
    "E": [0, 1],
    "W": [0, -1],
}


def in_bounds(grid, i, j):
    in_bound_v = i < len(grid) and i >= 0
    in_bounds = in_bound_v and j < len(grid[i]) and j >= 0
    return in_bounds


def init_traversal(grid, i, j):
    q = Queue()
    if in_bounds(grid, i - 1, j) and grid[i - 1][j] in {"F", "|", "7"}:
        q.put([i - 1, j, "N"])
    if in_bounds(grid, i + 1, j) and grid[i + 1][j] in {"J", "|", "L"}:
        q.put([i + 1, j, "S"])
    if in_bounds(grid, i, j + 1) and grid[i][j + 1] in {"J", "-", "7"}:
        q.put([i, j + 1, "E"])
    if in_bounds(grid, i, j - 1) and grid[i][j - 1] in {"L", "-", "F"}:
        q.put([i, j - 1, "W"])
    return q


def traverse_pipe(grid, q):
    max_steps = 0
    while not q.empty():
        pipes = q.qsize()
        max_steps += 1
        for i in range(pipes):
            pipe = q.get()
            i, j, dir = pipe[0], pipe[1], pipe[2]
            kind = grid[i][j]

            next_dir = new_dir[dir][kind] if kind in new_dir[dir] else dir

            next_pipe = [
                i + new_dir_inc[next_dir][0],
                j + new_dir_inc[next_dir][1],
                next_dir,
            ]
            next_pipe_kind = grid[next_pipe[0]][next_pipe[1]]
            if next_pipe_kind == ".":
                continue

            grid[i] = grid[i][:j] + "." + grid[i][j + 1 :]
            q.put(next_pipe)

    return max_steps


def main():
    f = open("input2.txt", "r")
    grid = f.readlines()

    i = -1
    j = -1
    found_animal = False
    for line_idx in range(len(grid)):
        line = grid[line_idx]
        for char_idx in range(len(line)):
            char = line[char_idx]
            if char == "S":
                i = line_idx
                j = char_idx
                found_animal = True
                grid[i] = grid[i][:j] + "." + grid[i][j + 1 :]
                break
        if found_animal:
            break
    q = init_traversal(grid, i, j)
    print("Max Steps: ", traverse_pipe(grid, q))


if __name__ == "__main__":
    main()
