def extract_digits(line):
    """Extract the first and last digit from a line."""
    digits = [char for char in line if char.isdigit()]
    if digits:
        return int(digits[0] + digits[-1])  # Combine first and last digit
    return 0  # If no digits, return 0

def calculate_sum(file_path):
    total_sum = 0
    with open(file_path, 'r') as file:
        for line in file:
            total_sum += extract_digits(line.strip())  # Process each line
    return total_sum

# Run the solution
file_path = "input.txt"  # Your input file
result = calculate_sum(file_path)
print(result)
