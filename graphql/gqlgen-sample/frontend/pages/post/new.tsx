import React, { useState } from 'react'
import { useMutation } from '@apollo/client'
import { CREATE_POST_M } from '../../graphql/mutations/createPosts'
import { Post } from '../../src/types/post'
import { Box, Button, TextField } from '@material-ui/core'

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
