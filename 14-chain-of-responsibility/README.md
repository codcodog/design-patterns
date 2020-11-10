责任链模式
===========

它包含了一些命令对象和一系列的处理对象.  
每一个处理对象决定它能处理哪些命令对象，它也知道如何将它不能处理的命令对象传递给该链中的下一个处理对象.

责任链模式是一条链，链上有多个节点，每个节点都有各自的责任.  
当有输入时，第一个责任节点看自己能否处理该输入，如果可以就处理。如果不能就交由下一个责任节点处理.

一般在处理多个 `if...else` 或者 `swich..case` 时适用.