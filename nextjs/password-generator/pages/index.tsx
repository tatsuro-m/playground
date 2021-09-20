import React, { useState } from 'react'
import { Box, Grid, Paper } from '@material-ui/core'
import { makeStyles } from '@material-ui/core'
import DisplayBlock from '../components/DisplayBlock'
import KindDropdown from '../components/KindDropdown'

const useStyles = makeStyles({
  background: {
    background: '#0a2d4d',
    height: 800,
    width: '100%',
  },
})

export const Home: React.VFC = () => {
  const classes = useStyles()
  const [password] = useState('hogehoge')
  const [kind, setKind] = useState('random')

  const handleKindChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    if (['random', 'simple', 'pin'].includes(event.target.value)) {
      setKind(event.target.value)
    }
  }

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
              <Grid item xs={4}>
                <Box m={2}>
                  <KindDropdown
                    kind={kind}
                    handleKindChange={handleKindChange}
                  />
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
