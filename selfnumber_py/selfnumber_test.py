import unittest

def sumOfSelfNumbers(_from, _to):
    notSelfNumbers = {d(i) for i in range(_to)}

    return sum(set(range(_from, _to)) - notSelfNumbers)

def d(n):
    return n + sumOfDigits(n)

def sumOfDigits(n):
    if n < 10:
        return n

    return sumOfDigits(n//10) + n % 10


class SelfNumberTest(unittest.TestCase):
    def test_sum_of_selfnumbers(self):
        self.assertEqual(1+3+5+7+9, sumOfSelfNumbers(1, 10))
        self.assertEqual(1227365, sumOfSelfNumbers(1, 5000))
        self.assertEqual(1222233933479, sumOfSelfNumbers(1, 5_000_000))

    def testGenerator(self):
        self.assertEqual(101, d(91))
        self.assertEqual(1003, d(1001))