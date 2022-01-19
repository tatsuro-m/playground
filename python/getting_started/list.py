arr = ["a", "b", "c", "d", "e"]
print(arr[0])
print(arr[-1])

# 範囲参照。
# Go と同じで最初のインデックスは含むが最後のインデックスは含まずその１個手前まで。
print(arr[0:2])
print(arr[1:4])

arr.insert(5, "f")
arr.remove("f")
print(arr)
print(len(arr))
