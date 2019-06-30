package main

/*
	条件变量，经常与锁结合使用

	条件变量结构体：
		type Cond struct {
			noCopy noCopy

			// L is held while observing or changing the condition
			L Locker

			notify  notifyList
			checker copyChecker
		}

	条件变量的操作函数：
		Wait方法：
			1、阻塞
			2、释放自己拥有的互斥锁(cond.L.Unlock())，这2步为一个原子操作
			3、当被唤醒后，Wait方法返回时，解除阻塞并重新获取互斥锁(cond.L.Lock())
		Signal方法：
			给一个阻塞在该条件变量上的Go程发送唤醒通知
		Broadcast方法：
			给阻塞在该条件变量上的所有Go程发送唤醒通知
*/
func main() {

}
