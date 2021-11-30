import React, { useEffect, useReducer } from "react";
import { AppProps } from "next/app";
import Head from "next/head";
import { CssBaseline } from "@material-ui/core";
import { ThemeProvider } from "@material-ui/core/styles";
import authReducer from "../lib/authReducer";
import { listenAuthState } from "../lib/firebase";
import theme from "../theme";
import { ApolloProvider } from "@apollo/client";
import AuthContext from "../lib/AuthContext";
import client from "../apollo-client";

export default function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector("#jss-server-side");
    jssStyles?.parentElement?.removeChild(jssStyles);
  }, []);

  const [state, dispatch] = useReducer(
    authReducer.reducer,
    authReducer.initialState
  );
  useEffect(() => {
    return listenAuthState(dispatch);
  }, []);

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
          <ApolloProvider client={client}>
            <Component {...pageProps} />
          </ApolloProvider>
        </AuthContext.Provider>
      </ThemeProvider>
    </>
  );
}
