import unittest
import unittest.mock as Mock

mock = Mock.MagicMock()
mock.return_value = 5
print(mock())

