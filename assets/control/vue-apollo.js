import { ApolloClient } from 'apollo-client'
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'

// HTTP connection to the API
const httpLink = createHttpLink( {

    // uri: '192.168.67.2',
    // uri: 'http://192.168.67.199:5555/query',
    uri: 'https://suds.ictgroup.mn/query',
    // uri: '/query',
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
