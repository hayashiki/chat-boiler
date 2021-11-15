import React from 'react'
import { AppProps } from 'next/app'
import Head from 'next/head'
import { useApollo } from "@/lib/graphql";
import { ApolloProvider } from "@apollo/client";

function App({ Component, pageProps }: AppProps): JSX.Element {
  React.useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    if (jssStyles) {
      jssStyles.parentElement?.removeChild(jssStyles)
    }
  }, [])

  const apolloClient = useApollo(pageProps.initialApolloState)

  return (
    <React.Fragment>
      <Head>
        <title>hayashiki | scaffold </title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <ApolloProvider client={apolloClient}>
        <Component {...pageProps} />
      </ApolloProvider>
    </React.Fragment>
  )
}

export default App

