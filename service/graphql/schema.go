package graphql

var Schema = `
    schema {
        query: Query
        mutation: Mutation
    }
    type Query {
        product(id: ID!): Product
        searchProduct(title: String!): ID
    }
    type Mutation {
        createProduct(title: String!, price: Float!, manufacturer: String!, description: String): ID
        updateProduct(id: ID!, price: Float!): ID
        deleteProduct(id: ID!): ID
    }
    type Product {
        id: ID!
        catID: ID!
        title: String!
        price: Float!
        manufacturer: String!
        description: String
        created: String!
        updated: String!
    }
`
