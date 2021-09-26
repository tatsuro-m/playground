import client from '../apollo-client'
import { gql } from '@apollo/client'
import { GetStaticProps } from 'next'
import React from 'react'

export const getStaticProps: GetStaticProps = async () => {
  const { data } = await client.query({
    query: gql`
      query Posts {
        posts {
          id
          title
          createdAt
          updatedAt
        }
      }
    `,
  })

  return {
    props: {
      posts: data.posts,
    },
  }
}

interface Post {
  id: number
  title: string
  createdAt: string
  updatedAt: string
}

interface Props {
  posts: Post[]
}

export const Home: React.VFC<Props> = (props) => {
  return (
    <>
      <h1>Posts</h1>
      <h1>{props.posts[0].title}</h1>
      <h1>{props.posts[1].id}</h1>
      <h1>{props.posts[1].createdAt}</h1>
      <h1>{props.posts[1].updatedAt}</h1>
    </>
  )
}

export default Home
