import React from 'react'
import { Grid, Input, makeStyles, Slider } from '@material-ui/core'

interface Props {
  length: number
  handleSliderChange: (event: any, newValue: number) => void
  handleInputChange: (event: React.ChangeEvent<HTMLInputElement>) => void
  handleBlur: () => void
}

const useStyles = makeStyles({
  input: {
    width: 42,
  },
})

const LengthSlider: React.VFC<Props> = (props) => {
  const classes = useStyles()
  const { length, handleBlur, handleInputChange, handleSliderChange } = props

  return (
    <>
      <Grid container spacing={2} alignItems="center">
        <Grid item>
          <p>長さ</p>
        </Grid>
        <Grid item xs>
          <Slider
            value={typeof length === 'number' ? length : 0}
            onChange={handleSliderChange}
            aria-labelledby="input-slider"
          />
        </Grid>
        <Grid item>
          <Input
            className={classes.input}
            value={length}
            margin="dense"
            onChange={handleInputChange}
            onBlur={handleBlur}
            inputProps={{
              step: 1,
              min: 0,
              max: 300,
              type: 'number',
              'aria-labelledby': 'input-slider',
            }}
          />
        </Grid>
      </Grid>
    </>
  )
}

export default LengthSlider
