桥接模式
==========

把事物对象和其具体行为、具体特征分离开来，使它们可以各自独立的变化.

例如：圆形，三角形归于抽象的形状之下，而画圆，画三角归于实现行为的画图之下，然后由形状调用画图.

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换.