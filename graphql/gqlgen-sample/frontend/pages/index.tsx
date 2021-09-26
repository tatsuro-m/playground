import client from '../apollo-client'
import { gql } from '@apollo/client'
import { GetStaticProps } from 'next'

export const getStaticProps: GetStaticProps = async () => {
  const { data } = await client.query({
    query: gql`
      query Posts {
        posts {
          id
          title
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

export const Home = ({ posts }): JSX.Element => {
  return (
    <>
      <h1>Posts</h1>
      <h1>{posts[0].title}</h1>
    </>
  )
}

export default Home
