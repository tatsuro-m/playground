import React from 'react'
import { Paper } from '@material-ui/core'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  background: {
    background: '#ffffff',
    height: 150,
    width: '100%',
  },
})

interface Props {
  password: string
}

const DisplayBlock: React.VFC<Props> = (props) => {
  const classes = useStyles()

  return (
    <Paper className={classes.background}>
      <p>{props.password}</p>
    </Paper>
  )
}

export default DisplayBlock
