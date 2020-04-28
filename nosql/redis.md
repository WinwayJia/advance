key 淘汰机制
单线程

内存暴涨-rehash


list,hash,zset数据量小时都使用ziplist，无法使用时，list使用list链表
set使用intset, 如法使用时，转换成dict

zset使用dict和skiplist
- dict中只存score
- skiplist中存field key和score
- zadd时field key长度或者field长度超了，会调用zsetConvert从ziplist转到dick+skiplist

