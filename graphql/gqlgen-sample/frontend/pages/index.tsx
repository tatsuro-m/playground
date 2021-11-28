import { gql, useQuery } from '@apollo/client'
import React from 'react'
import {
  Box,
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@material-ui/core'
import { firebaseUser, onLogout } from '../src/lib/firebase'
import Link from 'next/link'
import { Posts } from '../src/types/post'

const POSTS_QUERY = gql`
  query GetPosts {
    posts {
      id
      title
      createdAt
      updatedAt
    }
  }
`

export const Home: React.VFC = () => {
  const { loading, error, data } = useQuery<Posts>(POSTS_QUERY)

  if (loading) return <p>Loading...</p>
  if (error) return <p>Error: {JSON.stringify(error)}</p>

  const { posts } = data

  return (
    <Box m={10}>
      {firebaseUser() ? (
        <Button variant="contained" onClick={() => onLogout()}>
          ログアウト
        </Button>
      ) : (
        <Link href="/sign_in">
          <Button variant="contained" color="primary">
            ログイン画面
          </Button>
        </Link>
      )}
      <p>
        {firebaseUser()
          ? firebaseUser().displayName + 'でログインしています'
          : 'ログインしていません'}
      </p>
      <Link href="/post/new">
        <a>投稿を作成する</a>
      </Link>
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
            {posts.map((post) => (
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
