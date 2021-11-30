import React, { useState } from "react";
import { Box, Button, TextField } from "@material-ui/core";
import { useCreatePostMutation } from "../../generated/graphql";

export const New: React.VFC = () => {
  const [createPost, { data, loading, error }] = useCreatePostMutation();
  console.log(JSON.stringify(error));

  const [input, setInput] = useState("");

  const updateInput = (event) => {
    console.log(event.target.value);
    setInput(event.target.value);
  };

  const post = data;
  console.log(post);

  return (
    <Box m={10}>
      <TextField
        id="outlined-basic"
        label="Outlined"
        variant="standard"
        value={input}
        onChange={(e) => updateInput(e)}
      />
      <Button
        variant="contained"
        color="primary"
        onClick={() =>
          createPost({
            variables: { title: input },
          })
        }
      >
        投稿を作成
      </Button>
      {loading ? <p>loading...</p> : <p>done</p>}
    </Box>
  );
};

export default New;
