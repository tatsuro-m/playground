import { ApolloClient, InMemoryCache } from '@apollo/client'
import { firebaseUser } from './lib/firebase'

const getJWT = async () => {
  return await firebaseUser().getIdToken(true)
}

const client = new ApolloClient({
  uri: process.env.NEXT_PUBLIC_GRAPHQL_SERVER_URI,
  cache: new InMemoryCache(),
  headers: {
    authorization: firebaseUser() ? `Bearer: ${getJWT()}` : '',
  },
})

export default client
