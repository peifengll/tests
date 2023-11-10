格式
每次提交，Commit message 包括三个部分：Header（标题），Body（内容详情） 和 Footer（结尾）。
其中 Header是必须的，Body 和 Footer可以省略。

<type>(<scope>): <subject>
// 空一行
<body>
// 空一行
<footer>  
1
2
3
4
5
Header部分
只有一行，包括两个字段：type（必需）和subject（必需）。

type
用于说明commit的类别，只允许有以下类别

feat: 新功能（feature）
fix: 修补bug
docs: 文档（documentation）
style: 格式（不影响代码运行的变动）
refactor: 重构（即不是新增功能，也不是修改bug的代码变动）
chore: 构建过程或辅助工具的变动
revert: 撤销，版本回退
perf: 性能优化
test：测试
improvement: 改进
build: 打包
ci: 持续集成

scope
用来说明本次Commit影响的范围，即简要说明修改会涉及的部分。选填。

subject
标题，简述本次操作

Body部分
选填。对上面subject里内容的展开，在此做更加详尽的描述，内容里应该包含修改动机和修改前后的对比。

Footer部分
选填。
Footer 部分只用于两种情况。
（1）不兼容变动
如果当前代码与上一个版本不兼容，则 Footer 部分以BREAKING CHANGE开头，后面是对变动的描述、以及变动理由和迁移方法。
（2）关闭 Issue
如果当前 commit 针对某个issue，那么可以在 Footer 部分关闭这个 issue 。
Closes #234 也可以一次关闭多个 issue 。
Closes #123, #245, #992

Revert
还有一种特殊情况，如果当前 commit 用于撤销以前的 commit，则必须以revert:开头，后面跟着被撤销 Commit 的 Header。
revert: feat(订单模块): 回退当前版本2到1

例
git add .
git status
git commit“feat: xxx「xx列表页」 功能开发
>
> -xxx
> -xxx”