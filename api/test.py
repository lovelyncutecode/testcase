import unittest
from converter import convert

class ConverterTest(unittest.TestCase):
	def testZero(self):
		self.assertEqual(convert(0,'UAH'), 0.0)

	def testLower(self):
		self.assertEqual(convert(0,'uah'), 0.0)	

	def testCurrency(self):
		self.assertEqual(convert(0,'UAH'), 0.0)
		self.assertEqual(convert(0,'EUR'), 0.0)
		self.assertEqual(convert(0,'GBP'), 0.0)

	def testError(self):
		with self.assertRaises(ValueError):
			convert(0,'RUB')
		with self.assertRaises(ValueError):
			convert('u','UAH')    	

if __name__ == '__main__':
    unittest.main()