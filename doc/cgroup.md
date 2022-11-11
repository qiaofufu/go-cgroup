# 1.介绍

## 1-1.术语

"cgroup"代表"control group",从不使用大写.单数形式用于指定整个特征,也用作修饰符如"cgroup controllers".当明确提到多个单独的控制组时,使用"cgroups"的复数形式.

## 1-2.什么是cgroup

- cgroup是一种机制.
- 用于分层组织进程,并以受控和可配置的方式沿着层次结构分配系统资源.

- cgroup很大程度上由两部分组成- `core`和`controllers`.cgroup core 主要负责分层组织进程.
- cgroup controller 通常负责沿着层次结构分配具体类型的系统资源,尽管有实用的控制器服务于资源分配以外的目的.

- cgroup形成树的结构并且每一个在系统中的进程属于并且仅能属于一个cgroup.
- 进程的所有线程属于同样的cgroup.
- 在创建时,所有的进程被放在父进程当时所属的cgroup中.
- 进程可以被迁移到另外一个cgroup.
- 进程的迁移并不影响已经存在后代的进程

- 遵循一定的结构性限制,控制器可以在cgroup中选择项地启用或禁用.
- 所有控制器的行为是阶级式的 - 如果控制器在cgroup中启用, 它影响所有属于这个cgroup包括子层次结构的cgroup的进程.
- 当控制器在嵌套的cgroup中被启用时,它总是进一步限制资源分配.
- 在层次结构离root近的限制设置不能被更远的覆盖.

# 2.基本操作

## 2-1. 挂载

与v1版本不同, cgroup v2 仅有一个层次结构. cgroup v2的层次结构能用下面的挂载命令挂载:

```bash
mount -t cgroup2 none $MOUNT_POINT
```

- cgroup2 文件系统有一个魔术数字`0x63677270`("cgrp").
- 所有支持v2且没有绑定到v1层次结构的控制器自动的绑定到v2 层次结构中并在根部显示出来.
- 在v2 层次结构中没有使用的控制器能绑定到其他层次结构.允许以完全向后兼容的方式将v2层次结构与遗留的v1多个层次结构混合.
- 只有控制器在当前层次结构不再被引用之后,控制器才能跨越层级移动.
- 因为每一个cgroup控制器的状态是异步销毁的而控制器可能有拖延的引用, 在上一个层次结构最终取消挂载之后控制器可能没有立即出现在v2层次结构上.
- 同样的,控制器应该被完全禁用才能移除unified层次结构并且它可能需要一些时间才能被其他层次结构使用.
- 此外,由于控制器内部的依赖,其他控制器可能需要被禁用.



- 虽然对开发和手动配置很有用,不推荐在生产环境下, 在v2和其他层次结构间动态移动控制器.
- 推荐在系统启动后控制器使用之前,决定层次结构和控制器之间的关联.

- 在向v2过渡的期间,在手动干预之前可能会发生: 系统管理软件仍然会自动挂载v1 cgroup 文件系统并且boot期间劫持所有的控制器.
- 为了让测试和实验变得更加容易, 允许设置内核参数`cgroup_no_v1=`来禁用v1控制器,并使他们在v2中始终可用

cgroup v2 现在支持下面的挂载选项

### nsdelegate

​	太抽象了没看懂!!!

```
Consider cgroup namespaces as delegation boundaries.  This
	option is system wide and can only be set on mount or modified
	through remount from the init namespace.  The mount option is
	ignored on non-init namespace mounts.  Please refer to the
	Delegation section for details.
```

## 2-2. 组织进程和线程

### 进程 Processes

最初, 只有root cgroup存在,所有的进程都属于root cgroup.

可以通过创建子文件夹创建子cgroup

```bash
mkdir $CGROUP_NAME
```

一个给定的cgroup可能有多个子cgroup形成一个树型结构.

每个cgroup有一个可读可写的接口文件`cgroup.procs`.

在读文件时,它列出所有属于这个cgroup的进程的Pids(每个一行).

Pids时无序的并且可能出现同样的PID不止一次,例如进程被移动到另外一个cgroup后再移动回来或pid在读的时候被收回.

通过写PID到目标cgroup的`cgroup.procs`文件,进程能够被移动到一个cgroup中.

在一次`write(2)`调用中只能将一个进程移动到cgroup.

如果一个进程由多个线程组成, 写任何一个线程的PID会迁移这个进程所有线程.

当一个进程fork一个子进程时,新的进程加入这次操作时父进程所属的cgroup.

退出之后,进程逗留在关联的退出时所属的cgroup中,直到被reaped; 然而僵尸进程不会出现在"cgroup.procs"里同时也不能移动到另外的cgroup.

一个没有任何儿子或活跃进程的cgroup可以通过删除文件夹销毁.

提示: 没有任何儿子和仅与僵尸进程相关联被认为是空的,可以被删除:

```bash
rmdir $CGROUP_NAME
```

"/proc/$PID/cgroup" 列出了进程是哪个cgroup的成员.

如果遗留的cgroup在系统中使用,这个文件可能包含多行. 每一行是一个层次结构.

cgroup v2的记录总是这样的格式"0::$PATH"

```bash
cat /proc/$PID/cgroup
...
0::/test-cgroup/test-cgroup-nested
```

如果进程变成僵尸进程并且它关联的cgroup随后被删除, "(deleted)"将被添加到path的后面

```bash
cat /proc/$PID/cgroup
...
0::/test-cgroup/test-cgroup-nested (deleted)
```

### 线程 Threads

- cgroup v2 支持控制器子集线程粒度, 去支持需要跨越进程的线程组分层资源分配的用例.
- 默认的,一个进程的所有线程属于同一个cgroup, 这个cgroup也作为资源域去承载非特定的进程或线程的资源消耗.
- 线程模式允许线程分布在整个子树中同时仍然为他们保持公共资源域
- 支持线程模式的控制器被称作线程控制器(`threaded controllers`).反之被称为域控制器(`domain controllers`).

- 标记cgroup 为线程 使它加入父资源域作为线程cgroup.
- 父资源域可能是另外一个线程cgroup,它的资源域是在层次结构中更上一层的.
- 线程树的根,即最近的非线程祖先,可呼唤的称为线程域或线程树的根并且作为整个子树的资源域
- 在线程树的内部, 一个进程的线程能被放到不同的cgroup里, 并且不受非内部进程限制.-在非叶cgroup中无论他们有没有线程,都可以启用线程控制器.
- 





