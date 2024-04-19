package database

import(
  "os"
  "log"
  "time"
  "context"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "github.com/sajagsubedi/ContactLync/graph/model"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
)
 var connectionString string=os.Getenv("MONGO_URI")
 
type DB struct {
  client *mongo.Client
}

func ConnectDb() *DB {
  client,
  err:= mongo.NewClient(options.Client().ApplyURI(connectionString))
  if err != nil {
    log.Fatal(err)
  }

  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal(err)
  }

  err = client.Ping(ctx, readpref.Primary())
  if err != nil {
    log.Fatal(err)
  }

  return &DB {
    client: client,
  }
}

//methods in db
func(db *DB) CreateUser(input *model.CreateUserInput) *model.User {
  return &model.User{}
}

func(db *DB) UpdateUser(input *model.UpdateUserInput)*model.User {
  return &model.User{}
}

func(db *DB) DeleteUser(id string) *model.User {
  return &model.User{}

}

func(db *DB) Users() []*model.User {
  return []*model.User{}

}

func(db *DB) User(id string) *model.User {
 userCollec:=db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  _id, _ :=primitive.ObjectIDFromHex(id)
  filter:=bson.M{"_id":_id}
  
  var foundUser model.User
  err:=userCollec.FindOne(ctx,filter).Decode(&foundUser)
  if err!=nil{
    log.Fatal(err)
  }
  return &foundUser
}

func(db *DB) UserByFilter(input *model.FilterInput) []*model.User {
  return []*model.User{}
}