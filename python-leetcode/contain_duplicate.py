nums = [1, 2, 3, 4]
a = set()
for num in nums: 
    if num in a:
        print("Yes")
    else:
        a.add(num)
print("No")
