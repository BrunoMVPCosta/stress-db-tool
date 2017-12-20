Stress db tool
==================

Create a configuration file named `conf.json` with the format

```
{
  "MongoConnectionString": "<mongo-connectionstring>", 
  "InsertMany": true,
  "BulkInsert": true,
  "BulkInsertWithSleep": true
}
```