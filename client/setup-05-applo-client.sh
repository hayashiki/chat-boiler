# 手作業
# _app.tsxにProviderはさみこむ
#import { useApollo } from "@/lib/graphql";
#import UserProvider from "@/context/useContext";
#
#function App({ Component, pageProps }: AppProps): JSX.Element {
#  const apolloClient = useApollo(pageProps.initialApolloState);
#
#  return (
#    <React.Fragment>
#      <Head>
#        <title>Mentions</title>
#        <meta
#          name="viewport"
#          content="minimum-scale=1, initial-scale=1, width=device-width"
#        />
#      </Head>
#      <ApolloProvider client={apolloClient}>
#        <UserProvider>
#          <Component {...pageProps} />
#        </UserProvider>
#      </ApolloProvider>

yarn add @apollo/react-hooks @apollo/client graphql
yarn add -D @graphql-codegen/cli @graphql-codegen/typescript @graphql-codegen/typescript-operations @graphql-codegen/typescript-react-apollo apollo-cache-inmemory graphql-anywhere

# すでにあるのでは
mkdir -p src/lib
touch src/lib/graphql.ts

cat << EOS > src/lib/graphql.ts
import { useMemo } from "react";
import { ApolloClient, HttpLink, InMemoryCache, NormalizedCacheObject} from "@apollo/client";

let apolloClient: ApolloClient<NormalizedCacheObject>;
const isBrowser = typeof window === 'undefined';

function createApolloClient() {
  try {
    return new ApolloClient({
      ssrMode: !isBrowser,
      link: new HttpLink({
        uri: "/api/graphql/",
      }),
      cache: new InMemoryCache(),
    });
  } catch (error) {
    console.error("fail to create apollo client", error)
  }
}

export const initializeApollo = (
  initialState: NormalizedCacheObject | null = null
): ApolloClient<NormalizedCacheObject> => {
  const _apolloClient = apolloClient ?? createApolloClient();

  if (initialState) {
    _apolloClient.cache.restore(initialState);
  }
  if (!isBrowser) return _apolloClient;
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
};

export const useApollo = (
  initialState: NormalizedCacheObject
): ApolloClient<NormalizedCacheObject> => {
  const store = useMemo(() => initializeApollo(initialState), [initialState]);
  return store;
};

EOS

# codegen.yml

cat << EOS > codegen.yml
schema: '../schema/*.graphql'
documents:
  - ./src/**/*.graphql
  - ./src/**/*.gql
overwrite: true
generates:
  ./src/generated/graphql.ts:
    plugins:
      - typescript
      - typescript-operations
      - typescript-react-apollo
    config:
      skipTypename: false
      withHooks: true
      withHOC: false
      withComponent: false
EOS
