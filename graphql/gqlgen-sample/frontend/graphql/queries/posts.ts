import { gql } from '@apollo/client'

export const POSTS_QUERY = gql`
  query GetPosts {
    posts {
      id
      title
      createdAt
      updatedAt
    }
  }
`
