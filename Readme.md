### ContactLync

ContactLync is a GraphQL API built in Go for managing contact information. It allows you to perform CRUD operations on user data stored in MongoDB. Below is a guide to help you run the application, along with examples of GraphQL queries and mutations.

#### Features:
- Create, read, update, and delete user data
- Retrieve users based on specific filters

#### Setup:
1. Clone the repository:
   ```
   git clone https://github.com/sajagsubedi/ContactLync.git
   ```

2. Navigate to the project directory:
   ```
   cd ContactLync
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Set environment variables:
   - `PORT`: Port number for running the server
   - `MONGO_URI`: MongoDB connection URI

5. Build and run the application:
   ```
   go build
   ./ContactLync
   ```

#### Examples:

##### Query:
1. Retrieve all users:
   ```graphql
   query {
     users {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

2. Retrieve a user by ID:
   ```graphql
   query {
     user(id: "user_id") {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

3. Retrieve users by filter:
   ```graphql
   query {
     userByFilter(input: { field: "name", value: "John" }) {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

##### Mutation:
1. Create a new user:
   ```graphql
   mutation {
     createUser(input: { name: "Alice", phone: "1234567890", address: "123 Main St", email: "alice@example.com", relation: "Friend" }) {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

2. Update an existing user:
   ```graphql
   mutation {
     updateUser(input: { _id: "user_id", name: "Alice Smith" }) {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

3. Delete a user:
   ```graphql
   mutation {
     deleteUser(id: "user_id") {
       _id
       name
       phone
       address
       email
       relation
     }
   }
   ```

Feel free to explore and customize ContactLync according to your needs!
