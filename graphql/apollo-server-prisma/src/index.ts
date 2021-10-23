import {ApolloServer, gql} from 'apollo-server';
import {loadSchemaSync} from "@graphql-tools/load";
import {GraphQLFileLoader} from "@graphql-tools/graphql-file-loader";
import {join} from 'path';
import {addResolversToSchema} from "@graphql-tools/schema";

const schema = loadSchemaSync(join(__dirname, '../schema.graphql'), {
  loaders: [new GraphQLFileLoader()],
})

const books = [
  {
    title: 'The Awakening',
    author: 'Kate Chopin',
  },
  {
    title: 'City of Glass',
    author: 'Paul Auster',
  },
];

const resolvers = {
  Query: {
    books: () => books,
  },
};

const schemaWithResolvers = addResolversToSchema({schema, resolvers});
const server = new ApolloServer({schema: schemaWithResolvers});

server.listen().then(({url}) => {
  console.log(`🚀  Server ready at ${url}`);
});
