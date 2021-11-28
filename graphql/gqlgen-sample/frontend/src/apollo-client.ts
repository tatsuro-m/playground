import { ApolloClient, concat, HttpLink, InMemoryCache } from "@apollo/client";
import { firebaseUser } from "./lib/firebase";
import { setContext } from "@apollo/client/link/context";

const httpLink = new HttpLink({
  uri: process.env.NEXT_PUBLIC_GRAPHQL_SERVER_URI,
});

const authLink = setContext(async (_, { headers }) => {
  const token = firebaseUser() ? await firebaseUser().getIdToken(true) : null;

  if (token !== null) {
    return {
      headers: {
        ...headers,
        Authorization: `Bearer ${token}`,
      },
    };
  }
});

const client = new ApolloClient({
  cache: new InMemoryCache(),
  credentials: "include",
  link: concat(authLink, httpLink),
});

export default client;
