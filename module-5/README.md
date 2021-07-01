# Task

Driver for blob storage, the driver allows to store slice of bytes under some string key.

Driver should support general CRUD operations Create, Read, Update, Delete.
Interface example:
```go
type BlobStorage interface {
    Get(string) ([]byte, error)
    Put(string, []byte) (error)
    Delete(string) (error)
}
```
Driver interface  should have two implementations
  first implementation will just keep things in-memory, map data structure can be used.
  second implementation will store blobs on file-system, package (io)[https://golang.org/pkg/io/] and (os)[https://golang.org/pkg/os/] will be useful.


File structure:
{module}/blob/blob.go
{module}/blob/fs/fs.go
{module}/blob/memory/memory.go


