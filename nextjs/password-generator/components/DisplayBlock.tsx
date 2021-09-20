import React from 'react'
import { Paper } from '@material-ui/core'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  background: {
    background: '#ffffff',
    height: 200,
    width: '80%',
  }
})

const DisplayBlock: React.VFC = () => {
  const classes = useStyles()

  return <Paper className={classes.background}>ここに実際のパスワードが表示されるんだよ！</Paper>
}

export default DisplayBlock
