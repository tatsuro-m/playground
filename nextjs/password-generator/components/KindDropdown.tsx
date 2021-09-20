import React from 'react'
import { makeStyles, MenuItem, Select } from '@material-ui/core'

interface Props {
  kind: string
  handleKindChange: (event: React.ChangeEvent<HTMLSelectElement>) => void
}

const useStyles = makeStyles({
  dropdown: {
    background: '#0f4170',
    color: '#ffffff',
    width: '20em',
  },
})

const KindDropdown: React.VFC<Props> = (props) => {
  const classes = useStyles()

  return (
    <Select
      labelId="demo-simple-select-label"
      id="demo-simple-select"
      value={props.kind}
      onChange={props.handleKindChange}
      className={classes.dropdown}
      variant="outlined"
    >
      <MenuItem value="random">ランダムなパスワード</MenuItem>
      <MenuItem value="simple">覚えやすいパスワード</MenuItem>
      <MenuItem value="pin">暗証番号</MenuItem>
    </Select>
  )
}

export default KindDropdown
