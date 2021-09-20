import React, { useState } from 'react'
import {
  Box,
  Button,
  Checkbox,
  FormControlLabel,
  Grid,
  makeStyles,
  Paper,
} from '@material-ui/core'
import DisplayBlock from '../components/DisplayBlock'
import KindDropdown from '../components/KindDropdown'
import CachedIcon from '@material-ui/icons/Cached'
import CopyButton from '../components/CopyButton'
import LengthSlider from '../components/LengthSlider'

const useStyles = makeStyles({
  background: {
    background: '#0a2d4d',
    height: 800,
    width: '100%',
  },
})

export const Home: React.VFC = () => {
  const classes = useStyles()
  const [password] = useState('hcaoighaonvalghalnaogaoge')
  const [kind, setKind] = useState('random')

  const handleKindChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    if (['random', 'simple', 'pin'].includes(event.target.value)) {
      setKind(event.target.value)
    }
  }

  const [length, setLength] = useState(12)
  const handleSliderChange = (event: any, newValue: number) => {
    setLength(newValue)
  }

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLength(Number(event.target.value))
  }

  const handleBlur = () => {
    if (length < 0) {
      setLength(0)
    } else if (length > 300) {
      setLength(300)
    }
  }

  const [isInclude, setIsInclude] = useState({
    number: true,
    symbol: false,
  })
  const handleChecked = (event: React.ChangeEvent<HTMLInputElement>) => {
    setIsInclude({
      ...isInclude,
      [event.target.name]: event.target.checked,
    })
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
                      <CopyButton password={password} />
                    </Box>
                  </Grid>
                  <Grid item xs={12}>
                    <Paper>
                      <Grid container>
                        <Grid item xs={1} />
                        <Grid item xs={5}>
                          <LengthSlider
                            length={length}
                            handleSliderChange={handleSliderChange}
                            handleInputChange={handleInputChange}
                            handleBlur={handleBlur}
                          />
                        </Grid>
                        <Grid item xs={3}>
                          <FormControlLabel
                            control={
                              <Checkbox
                                checked={isInclude.number}
                                onChange={handleChecked}
                                name="number"
                                color="primary"
                              />
                            }
                            label="数字"
                          />
                        </Grid>
                        <Grid item xs={3}>
                          <FormControlLabel
                            control={
                              <Checkbox
                                checked={isInclude.symbol}
                                onChange={handleChecked}
                                name="symbol"
                                color="primary"
                              />
                            }
                            label="記号"
                          />
                        </Grid>
                      </Grid>
                    </Paper>
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
