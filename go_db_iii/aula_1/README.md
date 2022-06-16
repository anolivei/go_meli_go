# Exercício 1: SQL

1.
```sql
db.getCollection('restaurants').find({},{ restaurant_id : 1, name : 1, borough : 1,  cuisine : 1,  _id : 0 })
```

2.
```sql
db.getCollection('restaurants').find({name : /Bake/},{ restaurant_id : 1, name: 1, borough: 1,  cuisine : 1,  _id : 0 }).limit(3)
```

3.
```sql
db.getCollection('restaurants').find({$or :  [{cuisine : "Chinese"}, {cuisine : "Thai"}]}).count()

db.getCollection('restaurants').find({cuisine : {$in : ["Chinese", "Thai"]}}).count()
```

# Exercício 2: NoSQL

1.
```sql
db.getCollection('restaurants').find({ grades: { $elemMatch: { grade : "A", score : { $gt : 20 } } } } ).limit(3)
```

2.
```sql
db.getCollection('restaurants').find({"address.coord": {$size : 0}}).count()

db.restaurants.count( { "address.coord" : { $size : 0 } } )
```

3.
```sql
db.getCollection('restaurants').find({},{ name: 1, borough: 1,  cuisine : 1,  grades : {$slice : -1}, _id : 0 }).limit(3)
```

# Exercício 3:

1.
```sql
db.restaurants.aggregate([{"$group" : {_id:{cuisine:"$cuisine"}, count:{$sum:1}}},{$sort:{"count":-1}},{$limit:3}])
```

2.

3.


