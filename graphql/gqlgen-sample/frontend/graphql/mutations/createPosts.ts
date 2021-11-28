import { gql } from '@apollo/client'

export const CREATE_POST_M = gql`
  mutation createPost($title: String!) {
    createPost(input: { title: $title }) {
      id
      title
      createdAt
      updatedAt
    }
  }
`
