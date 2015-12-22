# golang tricks

### Conversions Between Strings and Byte Slices

Go does have a couple of optimizations for []byte to string and string to []byte conversions to avoid extra allocations (with more optimizations on the todo list).

The first optimization avoids extra allocations when []byte keys are used to lookup entries in map[string] collections: m[string(key)].

The second optimization avoids extra allocations in for range clauses where strings are converted to []byte: for i,v := range []byte(str) {...}.


### nil channel

Send and receive operations on a nil channel block forver. 


### range

The data values generated in the "range" clause are copies of the actual collection elements. They are not references to the original items. This means that updating the values will not change the original data. It also means that taking the address of the values will not give you pointers to the original data.
