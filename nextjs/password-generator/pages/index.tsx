import React, { useState } from 'react'
import { Box, Button, CssBaseline, Grid, Paper } from '@material-ui/core'
import { makeStyles } from '@material-ui/core'
import DisplayBlock from '../components/DisplayBlock'
import KindDropdown from '../components/KindDropdown'
import CachedIcon from '@material-ui/icons/Cached'

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

      <Grid container>
        <Grid item xs={1} />
        <Grid item xs={10}>
          <Paper elevation={6} className={classes.background}>
            <Grid container justifyContent="center">
              <Grid item xs={10}>
                <Grid container justifyContent="center">
                  <Grid item xs={12}>
                    <Box m={2}>
                      <DisplayBlock password={password} />
                    </Box>
                  </Grid>
                  <Grid item xs={5}>
                    <Box m={2}>
                      <KindDropdown
                        kind={kind}
                        handleKindChange={handleKindChange}
                      />
                    </Box>
                  </Grid>
                  <Grid item xs={2}>
                    <Button>
                      <CachedIcon fontSize="large" color="primary" />
                    </Button>
                  </Grid>
                  <Grid item xs={5}>
                    <Box m={2}>
                      <Button color="secondary">
                        安全なパスワードをコピー
                      </Button>
                    </Box>
                  </Grid>
                </Grid>
              </Grid>
            </Grid>
          </Paper>
        </Grid>
        <Grid item xs={1} />
      </Grid>

      <Box mb={5} />
    </>
  )
}

export default Home
