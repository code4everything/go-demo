## 文件名

文件名通常由单个或多个单词组成，用下划线 `_` 分隔

## 可见性

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）

## 一般结构

- 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。

- 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。

- 如果当前包是 main 包，则定义 main 函数。

- 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。