import React, { useState } from 'react'
import { gql, useMutation } from '@apollo/client'
import { Post } from '../../src/types/post'
import { Box, Button, TextField } from '@material-ui/core'

const CREATE_POST_M = gql`
  mutation createPost($title: String!) {
    createPost(input: { title: $title }) {
      id
      title
      createdAt
      updatedAt
    }
  }
`

export const New: React.VFC = () => {
  const [createPost, { data, loading, error }] =
    useMutation<Post>(CREATE_POST_M)
  console.log(JSON.stringify(error))

  const [input, setInput] = useState({
    title: '',
  })

  const updateInput = (event) => {
    console.log(event.target.value)
    setInput(event.target.value)
  }

  const post = data
  console.log(post)

  return (
    <Box m={10}>
      <TextField
        id="outlined-basic"
        label="Outlined"
        variant="standard"
        value={input.title}
        onChange={updateInput}
      />
      <Button
        variant="contained"
        color="primary"
        onClick={() => createPost({ variables: { title: '一旦決め打ち' } })}
      >
        投稿を作成
      </Button>
      {loading ? <p>loading...</p> : <p>done</p>}
    </Box>
  )
}

export default New
