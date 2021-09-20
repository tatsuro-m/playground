import React, { useState } from 'react'
import { Box, Grid, Paper } from '@material-ui/core'
import { makeStyles } from '@material-ui/core'
import DisplayBlock from '../components/DisplayBlock'

const useStyles = makeStyles({
  background: {
    background: '#0a2d4d',
    height: 800,
    width: '100%',
  },
})

export const Home: React.VFC = () => {
  const classes = useStyles()
  const [password, changePassword] = useState('hogehoge')
  console.log(changePassword)

  return (
    <>
      <Grid container justifyContent="center">
        <h1>Password Generator!</h1>
      </Grid>

      <Grid container justifyContent="center">
        <Grid item xs={10}>
          <Paper elevation={6} className={classes.background}>
            <Grid container justifyContent="center">
              <Grid item xs={10}>
                <Box m={8}>
                  <DisplayBlock password={password} />
                </Box>
              </Grid>
            </Grid>
          </Paper>
        </Grid>
      </Grid>

      <Box mb={5} />
    </>
  )
}

export default Home
