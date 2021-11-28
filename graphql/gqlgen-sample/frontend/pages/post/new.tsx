import React from 'react'
import { useMutation } from '@apollo/client'
import { CREATE_POST_M } from '../../graphql/mutations/createPosts'
import { Post } from '../../src/types/post'
import { Box, Button } from '@material-ui/core'

export const New: React.VFC = () => {
  const [createPost, { data, loading, error }] =
    useMutation<Post>(CREATE_POST_M)
  if (loading) return <p>Loading...</p>
  if (error) return <p>Error: {JSON.stringify(error)}</p>

  const post = data
  console.log(post)

  return (
    <Box m={10}>
      <Button variant="contained" color="primary" onClick={() => createPost}>
        投稿を作成
      </Button>
    </Box>
  )
}

export default New
