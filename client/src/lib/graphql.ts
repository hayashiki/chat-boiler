import { useMemo } from "react";
import { ApolloClient, HttpLink, split, InMemoryCache, NormalizedCacheObject} from "@apollo/client";
import { WebSocketLink } from '@apollo/client/link/ws';
import { getMainDefinition } from '@apollo/client/utilities';
import { setContext } from "@apollo/client/link/context";

let apolloClient: ApolloClient<NormalizedCacheObject>;
const isBrowser = typeof window === 'undefined';

const authLink = setContext((_, { headers }) => {
  // const client = getAuth0Client();
  //
  // if (!client) return { headers };
  // return client.getTokenSilently().then((token) => {
  //   return {
  //     headers: {
  //       ...headers,
  //       Authorization: token ? `Bearer ${token}` : '',
  //     },
  //   };
  // });
});

const wsLink =
  typeof window === 'undefined'
    ? null
    : new WebSocketLink({
  uri: `ws://localhost:8080/query`,
  options: {
    reconnect: true,
    lazy: true,
    // 認証は一旦とばす
    connectionParams: async () => {
      // const client = getAuth();
      // if (!client) return {};
      // const token = await client.getTokenSilently();
      // return {
      //   authToken: token,
      // }
    },
  }
});

const httpLink = new HttpLink({
  uri: 'http://localhost:8080/query'
});

const link =
  typeof window === 'undefined'
    ? authLink.concat(httpLink)
    : split(
    // split based on operation type
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === 'OperationDefinition' &&
        definition.operation === 'subscription'
      );
    },
    wsLink,
    authLink.concat(httpLink),
    );

function createApolloClient() {
  try {
    return new ApolloClient({
      ssrMode: !isBrowser,
      link: link,
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

