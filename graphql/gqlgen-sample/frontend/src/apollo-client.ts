import { ApolloClient, InMemoryCache } from '@apollo/client'
import { auth } from './lib/firebase'

const getJWT = async () => {
  return await auth.currentUser.getIdToken(true)
}

const client = new ApolloClient({
  uri: process.env.NEXT_PUBLIC_GRAPHQL_SERVER_URI,
  cache: new InMemoryCache(),
  headers: {
    // TODO https://www.apollographql.com/docs/react/networking/authentication/#reset-store-on-logout
    // ログアウト時にキャッシュを飛ばす必要があるかも
    authorization: auth.currentUser ? `Bearer: ${getJWT()}` : '',
  },
})

export default client
