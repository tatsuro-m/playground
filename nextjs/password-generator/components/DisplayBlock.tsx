import React from 'react'
import { Paper } from '@mui/material'
import { makeStyles } from '@mui/styles'

const useStyles = makeStyles({
  background: {
    background: '#ffffff',
    height: 200,
    width: '80%',
  }
})

const DisplayBlock: React.VFC = () => {
  const classes = useStyles()

  return <Paper className={classes.background}>ここに実際のパスワードが表示される</Paper>
}

export default DisplayBlock
