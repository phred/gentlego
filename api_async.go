package main



func testUsers() {
   userChan  := make(chan User) // make channels // HLinit
   storeChan := make(chan Store)

   go func() { userChan <- api.fetchUser() }()  // goroutines start immediately // HLspawn
   go func() { storeChan <- api.fetchStore() }()

   // this func does other work in parallel... // HLcollect

   user := <- userChan // and then waits for those operations to complete // HLcollect
   store := <- storeChan
   if (store.User != user.Id) {
      // ...
   }
}
