key 淘汰机制
单线程

内存暴涨-rehash


zset使用dict和skiplist
- dict中只存score
- skiplist中存field key和score
- zadd时field key长度或者field长度超了，会调用zsetConvert从ziplist转到dick+skiplist
