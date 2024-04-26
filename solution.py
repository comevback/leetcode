def main():
    num_cases = int(input())
    excute(num_cases)


def excute(num_cases, output=""):
    if num_cases == 0:
        print(output)
        return
    num_ints = int(input())
    ints = input()
    string_list = ints.split()
    ints_list = list(map(int, string_list))
    input_tuple = tuple(ints_list[:num_ints])
    result = str(sum_ints(input_tuple))
    output += result + "\n"
    return excute(num_cases - 1, output)

# 
def sum_ints(input_tuple):
    filtered_tuple = filter_tuple(input_tuple)
    return (sum_tuple(filtered_tuple))

# 
def sum_tuple(nums, total=0):
    if nums == ():
        return total
    return sum_tuple(nums[1:], total + nums[0] * nums[0])


def filter_tuple(input_tuple, index=0):
    if index >= len(input_tuple):
        return ()
    if input_tuple[index] > 0:
        return (input_tuple[index],) + filter_tuple(input_tuple, index + 1)
    else:
        return filter_tuple(input_tuple, index + 1)


if __name__ == "__main__":
    main()