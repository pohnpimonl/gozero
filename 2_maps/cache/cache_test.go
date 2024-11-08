package cache

// func TestCache(t *testing.T) {

// 	c := New()

// 	wg := sync.WaitGroup{}

// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()
// 		c.Set("a", "first")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		c.Set("b", "second")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		c.Set("c", "third")
// 	}()

// 	wg.Wait()

// }

// func TestCacheSync(t *testing.T) {
// 	c := NewSync()

// 	wg := sync.WaitGroup{}

// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()
// 		c.Set("a", "first")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		c.Set("b", "second")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		c.Set("c", "third")
// 	}()

// 	wg.Wait()
// }
