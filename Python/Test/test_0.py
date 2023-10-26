import pytest


@pytest.fixture(autouse=True)
def func():
    print("前置步骤")


class TestHaHa:

    # @pytest.mark.parametrize("a", [2, 3, 1])
    def test_01(self):
        """
        用例名称:test_01
        """
        assert 1 == 1, "test01 failed"

    # @pytest.mark.skip(reason="测试跳过")  # 设置跳过用例用于函数体外,可用于标注功能暂未实现的用例
    def test_02(self):
        """
        用例名称:test_02
        """
        assert 2 == 2, "test02 failed"

    # @pytest.mark.xfail(reason='测试标注失败')  # 该方法主要用到某些场景bug未修复或者用例存在问题可以直接标注为失败状态
    # def test_03(self):
    #     """
    #     用例名称:test_03
    #     """
    #     assert 3 == 3, "test03 failed"

    # # 用例执行过程中若条件不成立时跳出当前用例,用于函数体内
    # def test_04(self):
    #     """
    #     用例名称:test_04
    #     """
    #     if 2 > 1:
    #         pytest.skip("条件成立")
    #
    # condition = 'monica'
    #
    # @pytest.mark.skipif(condition == 'monica', reason='测试跳过')  # 该方法主要用到某些场景bug未修复或者用例存在问题可以直接标注为失败状态
    # def test_05(self):
    #     """
    #     用例名称:test_05
    #     """
    #     assert 3 == 3, "test05 failed"
    #
    # @pytest.mark.run(order=2)  # 设置用例执行顺序，数字越小优先执行
    # def test_06(self):
    #     """
    #     用例名称:test_06
    #     """
    #     assert 3 == 3, "test06 failed"
    #
    # @pytest.mark.run(order=1)  # 设置用例执行顺序，数字越小优先执行
    # def test_07(self):
    #     """
    #     用例名称:test_07
    #     """
    #     assert 3 == 3, "test07 failed"


if __name__ == '__main__':
    pytest.main()
