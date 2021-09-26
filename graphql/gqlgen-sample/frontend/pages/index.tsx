import client from '../apollo-client'
import { gql } from '@apollo/client'
import { GetStaticProps } from 'next'
import React from 'react'
import {
  Box,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@material-ui/core'

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
    <Box m={10}>
      <TableContainer component={Paper}>
        <Table aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>タイトル</TableCell>
              <TableCell align="right">作成日時</TableCell>
              <TableCell align="right">更新日時</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {props.posts.map((post) => (
              <TableRow key={post.id}>
                <TableCell component="th" scope="row">
                  {post.title}
                </TableCell>
                <TableCell align="right">{post.createdAt}</TableCell>
                <TableCell align="right">{post.updatedAt}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  )
}

export default Home
