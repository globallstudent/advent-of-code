from collections import Counter

def min_difference_sum(left, right):
    left.sort()
    right.sort()
    return sum(abs(a - b) for a, b in zip(left, right))

def read_input(file_path):
    left, right = [], []
    with open(file_path, 'r') as file:
        for line in file:
            a, b = map(int, line.split())
            left.append(a)
            right.append(b)
    return left, right  # Move return statement outside 'with'

    
def similarity_score(left, right):
    count_right = Counter(right)  # O(N)
    return sum(num * count_right.get(num, 0) for num in left)  # O(N)

file_path = "input.txt"
left_list, right_list = read_input(file_path)

result_1 = min_difference_sum(left_list, right_list) # Part 1
result_2 = similarity_score(left_list, right_list) # Part 2

print(result_1)
print(result_2)