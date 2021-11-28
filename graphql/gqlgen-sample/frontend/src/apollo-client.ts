import { ApolloClient, InMemoryCache } from '@apollo/client'
import { firebaseUser } from './lib/firebase'
import { cache } from 'browserslist'

export const getJWT = async () => {
  if (firebaseUser()) {
    const jwt = await firebaseUser().getIdToken(true)
    console.log(jwt)
    return jwt
  }
}

const createApolloClient = () => {
  return new ApolloClient({ uri: '', cache: new InMemoryCache() })
}

const client = new ApolloClient({
  uri: process.env.NEXT_PUBLIC_GRAPHQL_SERVER_URI,
  cache: new InMemoryCache(),
  credentials: 'include',
  headers: {
    // authorization: firebaseUser() ? `Bearer: ${getJWT()}` : '',
  },
})

export default client
