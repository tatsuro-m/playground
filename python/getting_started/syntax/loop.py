num = 1

while num < 10:
    print(num)
    num = num + 1

print('------------------------------')

# for
arr = [1, 2, 3, 4, 5]
for num in arr:
    print(num)

# while でも len を使えば同じことができる。
i = 0
while i < len(arr):
    print(arr[i])
    i = i + 1
