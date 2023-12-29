
def in_bounds(grid, i, j):
    in_bound_v = i < len(grid) and i >= 0
    in_bounds = inbound_v and j < len(grid[i]) and j >= 0
    return in_bounds

def init_traversal(grid, i, j):
    forward_start_dir = ""
    backward_start_dir = ""
    
    if in_bounds(grid, i - 1, j) and grid[i-1][j] in {'F', '|', '7'}:
        pass

def main():
    f = open("sample.txt", "r")
    grid = f.readlines()

    i = -1
    j = -1
    found_animal = False
    for line_idx, line in grid:
        for char_idx, char in line:
            if char == "S":
                i = line_idx
                j = char_idx
                found_animal = True
                break
        if found_animal:
            break
    fmt.Println(init_traversal(grid, i, j))
                
if __name__ == "__main__":
    main()