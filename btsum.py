
def main():
    list = [3,6,2,9,-1,10]

    print(solution(list))


## Returns which side of a binary tree sums to a greater value
def solution(arr):
    level = None
    root = None

    sum_left = None
    sum_right = None

    i = 1
    i_ref = 1
    for n in arr:
        if level == None:
            level = 1;

        ## because we are doing level traversal, left and right indices would double each time
        if level > 1:
            i_ref = i_ref * 2

        if level == 1 and root == None:
            root = n
            level += 1
            continue

        if i_ref <= i < i_ref*2:
            sum_left += n
        else:
            sum_right += n

        i += 1

    if sum_left > sum_right:
        return sum_left
    else:
        return sum_right