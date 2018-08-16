import pytest

def findSelfNumbers(to):
    selfNumberMap = {i:True for i in range(to)}

    for i in range(to):
        selfNumberMap[d(i)] = False

    selfnumbers = set({})

    return { i for i in range(to) if selfNumberMap[i] is True }
    
    return selfnumbers

def d(n):
    return n + sum(digits(n))

def digits(n):
    x = 1
    while True:
        if n < 10 ** x:
            return [n // 10 ** (x - i - 1) % 10 for i in range(x)]
        x += 1

def test_selfnumber():
    assert {1, 3, 5, 7, 9} == findSelfNumbers(10)
    assert 1227365 == sum(findSelfNumbers(5000))
    

def test_generate():
    assert 101 == d(91)
    assert 175+1+7+5 == d(175)

def test_digits():
    assert digits(80) == [8, 0]
    assert digits(91) == [9, 1]
    assert digits(101) == [1, 0, 1]
    assert digits(201) == [2, 0, 1]
    assert digits(2011) == [2, 0, 1, 1]