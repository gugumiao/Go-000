#### Week02 作业
##### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

应该Wrap这个error，交给上层做处理。dao层主要是做数据持久层的工作，负责与数据库进行联络的一些任务都封装在此，与业务强相关，
应该将错误Wrap上层，记录当时的堆栈和上下文信息(敏感信息过滤)。上层获取可以根据错误原因来决定是否处理。

