package main



func testUsers() {
   user := api.fetchUser(...) // blocking calls // HL
   store := api.fetchStore(...)

   // ... other work
   
   if (store.User != user.Id) {  // result used here // HL
      // ...
   }
}
