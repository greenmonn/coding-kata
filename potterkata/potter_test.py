import unittest

# 최저가를 구하자

def price(list):
    list.sort()
    list.reverse()

    seriesCount = getSeriesCount(list)

    if seriesCount == 5:
        sum1 = 8 * 5 * 0.75 + price(retrieveSeries(list, 5))
        sum2 = 8 * 4 * 0.8 + price(retrieveSeries(list, 4))
        return min(sum1, sum2)

    if seriesCount == 4:
        remain_list = retrieveSeries(list, 4)
        return 8 * 4 * 0.8 + price(remain_list)

    if seriesCount == 3:
        remain_list = retrieveSeries(list, 3)
        return 8 * 3 * 0.9 + price(remain_list)

    if seriesCount == 2:
        remain_list = retrieveSeries(list, 2)
        return 8 * 2 * 0.95 + price(remain_list)

    return sum(list) * 8

def retrieveSeries(list, limit):
    count = 0
    
    remain_list = [0, 0, 0, 0, 0]

    for i in range(len(list)):
        remain_list[i] = list[i]
        if list[i] != 0:
            count += 1
            if count > limit:
                return remain_list
            remain_list[i] -= 1
    return remain_list

def getSeriesCount(list):
    count = 0
    for i in list:
        if i != 0:
            count += 1
    return count


class PotterKataTest(unittest.TestCase):
    def test_retrieve_series(self):
        self.assertEqual([0, 0, 0, 0, 1], retrieveSeries([1, 1, 1, 1, 1], 4))

    def test_get_series_count(self):
        self.assertEqual(0, getSeriesCount([0, 0, 0, 0, 0]))
        self.assertEqual(2, getSeriesCount([1, 1, 0, 0, 0]))

    def test_simple(self):
        self.assertEqual(0, price([0, 0, 0, 0, 0]))
        self.assertEqual(8, price([1, 0, 0, 0, 0]))
        self.assertEqual(8, price([0, 1, 0, 0, 0]))
        self.assertEqual(8, price([0, 0, 1, 0, 0]))
        self.assertEqual(8, price([0, 0, 0, 1, 0]))
        self.assertEqual(8, price([0, 0, 0, 0, 1]))

    def test_simple_with_discounts(self):
        self.assertEqual(8 * 2 * 0.95, price([1, 1, 0, 0, 0]))
        self.assertEqual(8 * 3 * 0.9, price([1, 1, 1, 0, 0]))
        self.assertEqual(8 * 4 * 0.8, price([1, 1, 1, 1, 0]))
        self.assertEqual(8 * 5 * 0.75, price([1, 1, 1, 1, 1]))

    def test_several_discounts(self):
        self.assertEqual(8 + (8 * 2 * 0.95), price([2, 1, 0, 0, 0]))
        self.assertEqual(2 * (8 * 2 * 0.95), price([2, 2, 0, 0, 0]))
        self.assertEqual((8 * 4 * 0.8) + (8 * 2 * 0.95), price([2, 1, 2, 1, 0]))
        self.assertEqual(8 + (8 * 5 * 0.75), price([1, 2, 1, 1, 1]))

    def test_edge_cases(self):
        self.assertEqual(2 * (8 * 4 * 0.8), price([2, 2, 2, 1, 1]))
        self.assertEqual(3 * (8 * 5 * 0.75) + 2 * (8 * 4 * 0.8), price([5, 5, 4, 5, 4]))
