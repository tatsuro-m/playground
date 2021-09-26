import React, { useEffect, useReducer } from 'react'
import { AppProps } from 'next/app'
import Head from 'next/head'
import { CssBaseline } from '@material-ui/core'
import { ThemeProvider } from '@material-ui/core/styles'
import theme from '../src/theme'
import AuthContext from '../src/lib/AuthContext'
import authReducer from '../src/lib/authReducer'
import { listenAuthState } from '../src/lib/firebase'

export default function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    jssStyles?.parentElement?.removeChild(jssStyles)
  }, [])

  const [state, dispatch] = useReducer(
    authReducer.reducer,
    authReducer.initialState
  )
  useEffect(() => {
    return listenAuthState(dispatch)
  }, [])

  return (
    <>
      <Head>
        <title>graphql sample</title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <ThemeProvider theme={theme}>
        <AuthContext.Provider value={state}>
          <CssBaseline />
          <Component {...pageProps} />
        </AuthContext.Provider>
      </ThemeProvider>
    </>
  )
}
