import unittest
from hello import hello


class TestSayHello(unittest.TestCase):
    def test_sayHello(self):
        result = hello.SayHello()
        expected = "Hello World"
        self.assertEqual(result, expected, f"Should say {expected}")


if __name__ == '__main__':
    unittest.main()
