package database

import(
  "os"
  "fmt"
  "log"
  "time"
  "regexp"
  "context"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "github.com/sajagsubedi/ContactLync/graph/model"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
  "github.com/joho/godotenv"
)

type DB struct {
  client *mongo.Client
}

func ConnectDb() *DB {
  err:= godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error on loading .env file")
  }
  connectionString:= os.Getenv("MONGO_URI")
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
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  inserg,
  err:= userCollec.InsertOne(ctx, bson.M {
    "name": input.Name, "phone": input.Phone, "address": input.Address, "email": input.Email, "relation": input.Relation,
  })

  if err != nil {
    log.Fatal(err)
  }

  insertedID:= inserg.InsertedID.(primitive.ObjectID).Hex()
  newUser:= model.User {
    ID: insertedID,
    Name: input.Name,
    Phone: input.Phone,
    Address: input.Address,
    Email: input.Email,
    Relation: input.Relation,
  }
  return &newUser
}

func(db *DB) UpdateUser(input *model.UpdateUserInput)*model.User {
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  updateList:= bson.M {}

  if input.Name != nil {
    updateList["name"] = input.Name
  }

  if input.Email != nil {
    updateList["email"] = input.Email
  }
  if input.Phone != nil {
    updateList["phone"] = input.Phone
  }
  if input.Address != nil {
    updateList["address"] = input.Address
  }
  if input.Relation != nil {
    updateList["relation"] = input.Relation
  }
  _id,
  _:= primitive.ObjectIDFromHex(input.ID)

  filter:= bson.M {
    "_id": _id,
  }
  updateObj:= bson.M {
    "$set": updateList,
  }
  results:= userCollec.FindOneAndUpdate(ctx, filter, updateObj, options.FindOneAndUpdate().SetReturnDocument(1))
  var updatedUser model.User
  if err:= results.Decode(updatedUser); err != nil {
    log.Fatal(err)
  }
  return &updatedUser
}

func(db *DB) DeleteUser(id string) *model.DeleteUserResponse {
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  _id,
  _:= primitive.ObjectIDFromHex(id)
  filter:= bson.M {
    "_id": _id,
  }

  _,
  err:= userCollec.DeleteOne(ctx, filter)
  if err != nil {
    log.Fatal(err)
  }
  return &model.DeleteUserResponse {
    DeletedUserID: id,
  }

}

func(db *DB) Users() []*model.User {
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()

  var users []*model.User
  cursor,
  err:= userCollec.Find(ctx, bson.D {})
  if err != nil {
    log.Fatal(err)
  }
  if err = cursor.All(context.TODO(), &users); err != nil {
    log.Fatal(err)

  }
  return users
}

func(db *DB) User(id string) *model.User {
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  _id,
  _:= primitive.ObjectIDFromHex(id)
  filter:= bson.M {
    "_id": _id,
  }

  var foundUser model.User
  err:= userCollec.FindOne(ctx, filter).Decode(&foundUser)
  if err != nil {
    log.Fatal(err)
  }
  return &foundUser
}

func(db *DB) UserByFilter(input *model.FilterInput) []*model.User {
  userCollec:= db.client.Database("contactlync").Collection("users")
  ctx,
  cancel:= context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  regexPattern:= fmt.Sprintf(".*%s.*", input.Value)
  pattern,
  err:= regexp.Compile(regexPattern)
  if err != nil {
    log.Fatal(err)
  }
  filter:= bson.M {
    input.Field: bson.M {
      "$regex": primitive.Regex {
        Pattern: pattern.String(),
        Options: "i",
      }},
  }

  var users []*model.User
  cursor,
  err:= userCollec.Find(ctx, filter)
  if err != nil {
    log.Fatal(err)
  }
  if err = cursor.All(context.TODO(), &users); err != nil {
    log.Fatal(err)

  }
  return users
}