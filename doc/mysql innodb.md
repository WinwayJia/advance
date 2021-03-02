MySQL InnoDB

---

InnoDB两种行级锁：

| 锁类型           | 备注                       |
| ---------------- | -------------------------- |
| 共享锁（S Lock） | 允许事物读一行数据         |
| 排他锁（X Lock） | 允许事物删除活更新一行数据 |

|      | X      | S      |
| ---- | ------ | ------ |
| X    | 不兼容 | 不兼容 |
| S    | 不兼容 | 兼容   |

**意向锁**（Intention Lock）：为了支持在不同粒度上加锁，InnoDB支持的锁方式。意向锁是将锁定的对象分为多个层次，意向锁意味着事物希望在更细粒度（fine granularity）上加锁。InnoDB的意向锁即为表级别的锁。

| 意向锁类型 | 备注                               |
| ---------- | ---------------------------------- |
| 意向共享锁 | 事务想要获得一张表中某几行的共享锁 |
| 意向排他锁 | 事务想要获得一张表中某几行的排他锁 |

由于InnoDB存储引擎支持的是行级别的锁，因此意向锁其实不会阻塞除全表扫意外的任何请求。

|      | IS     | IX     | S      | X      |
| ---- | ------ | ------ | ------ | ------ |
| IS   | 兼容   | 兼容   | 兼容   | 不兼容 |
| IX   | 兼容   | 兼容   | 不兼容 | 不兼容 |
| S    | 兼容   | 不兼容 | 兼容   | 不兼容 |
| X    | 不兼容 | 不兼容 | 不兼容 | 不兼容 |



```sql
show engine innodb status \G
------------------------------
begin;
select name, age from user where uid=2 for update;
```

**表：information_schema.innodb_trx结构说明，reference： MySQL document**

该表只是显示了当前正在运行的InnoDB事务，并不能直接判断锁的一些情况。如果需要查看锁，则需要访问information_schema.innodb_locks

---

|      | 字段名                     | 说明                                                         |
| ---- | :------------------------- | ------------------------------------------------------------ |
| 1    | trx_id                     | InnoDB存储引擎内部唯一的事务ID                               |
| 2    | trx_state                  | 事务执行的状态<br />(RUNNING, LOCK WAIT, ROLLING BACK,  COMMITTING） |
| 3    | trx_started                | 事务开始的时间                                               |
| 4    | trx_requested_lock_id      | 正在等待的锁的ID                                             |
| 5    | trx_wait_started           | 事务等待开始时间                                             |
| 6    | trx_weight                 | 事务的权重，反映了一个事务修改和锁住的行数。在InnoDB存储引擎中，当发生死锁需要回滚时，InnoDB存储引擎会选择该值最小的进行回滚 |
| 7    | trx_mysql_thread_id        | MySQL中的线程ID，show processlist显示的结果                  |
| 8    | trx_query                  | 事务运行的SQL语句                                            |
| 9    | trx_operation_state        | 事务当前进行的操作                                           |
| 10   | trx_tables_in_use          | 被事务中正在执行SQL语句使用的表的数量                        |
| 11   | trx_tables_locked          | 被事务中正在执行SQL语句锁住的表的数量                        |
| 12   | trx_lock_structs           | 事务占有的锁的数量                                           |
| 13   | trx_lock_memory_bytes      | 事务中锁结构占有的内存大小                                   |
| 14   | trx_rows_locked            | 事务中锁住的行数（近似值）                                   |
|      | trx_rows_modified          | The number of modified and inserted rows in this transaction. |
|      | trx_concurrency_tickets    | A value indicating how much work the current transaction can do before being swapped out, as specified by the innodb_concurrency_tickets system variable. |
|      | trx_isolation_level        | The isolation level of the current transaction.（隔离级别）  |
|      | trx_unique_checks          | Whether unique checks are turned on or off for the current transaction. For example, they might be turned off during a bulk data load. |
|      | trx_foreign_key_checks     | Whether foreign key checks are turned on or off for the current transaction. For example, they might be turned off during a bulk data load. |
|      | trx_last_foreign_key_error | The detailed error message for the last foreign key error, if any; otherwise NULL. |
|      | trx_adaptive_hash_latched  | Whether the adaptive hash index is locked by the current transaction. When the adaptive hash index search system is partitioned, a single transaction does not lock the entire adaptive hash index. Adaptive hash index partitioning is controlled by innodb_adaptive_hash_index_parts, which is set to 8 by default. |
|      | trx_adaptive_hash_timeout  | Whether to relinquish the search latch immediately for the adaptive hash index, or reserve it across calls from MySQL. When there is no adaptive hash index contention, this value remains zero and statements reserve the latch until they finish. During times of contention, it counts down to zero, and statements release the latch immediately after each row lookup. When the adaptive hash index search system is partitioned (controlled by innodb_adaptive_hash_index_parts), the value remains 0. |
|      | trx_is_read_only           | A value of 1 indicates the transaction is read only.         |
|      | trx_autocommit_non_locking | A value of 1 indicates the transaction is a SELECT statement that does not use the FOR UPDATE or LOCK IN SHARED MODE clauses, and is executing with autocommit enabled so that the transaction contains only this one statement. When this column and TRX_IS_READ_ONLY are both 1, InnoDB optimizes the transaction to reduce the overhead associated with transactions that change table data. |

information_schema.innodb_locks

information_schema.innodb_lock_waits

#### 一致性非锁定读（consistent nonblocking read）： 

​	InnoDB存储引擎通过多版本控制的方式读取当前执行时间数据库中行的数据。这是默认的读取方式，即读取不会占用和等待表上的锁，但是在不同的事务隔离级别下，读取的方式不同，并不是在每个事务隔离级别下都是采用非锁定一致性读。此外，及时都是使用非锁定一致性读，但是对于快照的数据定义也各不相同。

​	一个行记录可能不止一个快照数据，一般成这个技术为行多版本技术。由此带来的并发控制，称为多版本并发控制（Multi Version Concurrency Control，MVVC）。

#### 一致性锁定读

```sql
SELECT col1, col2 FROM t WHERE id=1 LOCK IN SHARE MODE;
SELECT col1, col2 FROM t WHERE id=1 FOR UPDATE;
```

#### 自增长与锁

​	在InnoDB存储引擎的内存结构中，对每个含有自增长值的表都有一个自增长计数器（auto-increment counter）。这种锁其实是采用的一种特殊的表锁机制，为了提高插入性能，锁不是在一个事务完成后才释放，而是在完成对自增长值插入的SQL语句后立即释放。InnoDB存储引擎提供参数innodb_autoinc_lock_mode来控制自增长的模式，该参数默认值为1。另在InnoDB存储引擎中，自增长值必须是索引，同时必须是索引的第一个列。

### 锁的算法

#### 行锁的3中算法

| 类型          | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| Record Lock   | 单个行记录上的锁，总是会去锁住索引记录，如果InnoDB存储引擎表在建立的时候没有设立任何一个索引，那么InnoDB存储引擎会使用隐式的主见来进行锁定。 |
| Gap Lock      | 间隙锁，锁定一个范围，但不包含记录本身（左右都是开区间），Gap Lock的作用是为了阻止多个事务将记录插入到同一个范围内。 |
| Next-Key Lock | Gap Lock+Record Lock锁定一个范围，并且包含记录本身。<br />当查询的索引含有唯一属性时，InnoDB存储引擎会对Next-Key Lock进行优化，降级为Record Lock，即紧锁住索引本身，而不是范围。（左开右闭）<br />InnoDB使用Next-Key Locking的算法避免了幻读（Phantom Problem） |

**InnoDB存储引擎Next-Key Lock降级为Record Lock示例**

```sql
create table t(a int primary key);
insert into t select 1;
insert into t select 2;
insert into t select 5;
```

| 时间 | 会话A                                 | 会话B                           |
| ---- | ------------------------------------- | ------------------------------- |
| 1    | BEGIN;                                |                                 |
| 2    | SELECT * FROM t WHERE a=5 FOR UPDATE; |                                 |
| 3    |                                       | BEGIN;                          |
| 4    |                                       | INSERT INTO t SELECT 4;         |
| 5    |                                       | COMMIT;<br /># 成功，不需要等待 |
| 6    | COMMIT;                               |                                 |

```sql
CREATE TABLE z ( a INT, b INT, PRIMARY KEY(a), KEY(b) );
INSERT INTO z SELECT 1, 1;
INSERT INTO z SELECT 3, 1;
INSERT INTO z SELECT 5, 3;
INSERT INTO z SELECT 7, 6;
INSERT INTO z SELECT 10, 8;
```

||





---

#### Read Commited(Oracle, SqlServer) vs Repeatable Read(MySQL)

- *在RR隔离级别下，存在间隙锁，导致出现死锁的几率比RC大的多*
- *在RR隔离级别下，条件列未命中索引会锁表！而在RC隔离级别下，只锁行* 
- *在RC隔离级别下，**半一致性读(semi-consistent)**特性增加了update操作的并发性！*
  <u>在5.1.15的时候，innodb引入了一个概念叫做“semi-consistent”，减少了更新同一行记录时的冲突，减少锁等待。 所谓半一致性读就是，一个update语句，如果读到一行已经加锁的记录，此时InnoDB返回记录最近提交的版本，由MySQL上层判断此版本是否满足update的where条件。若满足(需要更新)，则MySQL会重新发起一次读操作，此时会读取行的最新版本(并加锁)！</u>

参考：

​	[互联网项目中mysql应该选什么事务隔离级别](https://zhuanlan.zhihu.com/p/59061106)