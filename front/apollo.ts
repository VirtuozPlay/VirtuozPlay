import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client/core';
import { createApolloProvider } from '@vue/apollo-option';

const cache = new InMemoryCache();

const link = new HttpLink({
    uri: '/graphql',
});

const apolloClient = new ApolloClient({
    cache,
    link,
});

const apolloProvider = createApolloProvider({
    defaultClient: apolloClient,
});

export default apolloProvider;
