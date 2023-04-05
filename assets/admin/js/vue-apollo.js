import { ApolloClient } from 'apollo-client'
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'

// HTTP connection to the API
const httpLink = createHttpLink({
    // uri: 'http://192.168.88.224:8282/query',
    // uri: 'http://103.87.255.139:443/query',
    // uri: 'http://192.168.88.224:8282/query',
    // uri: 'http://103.87.255.139:443/query', //server IP
    uri: 'http://192.168.43.96:8282/query', // aagii ID tvr
})

// Cache implementation
const cache = new InMemoryCache({
    addTypename: false
})

// Create the apollo client
export const apolloClient = new ApolloClient({
    link: httpLink,
    cache,
})
