import React from 'react'
import { MenuItem, Select } from '@material-ui/core'

interface Props {
  kind: string
  handleKindChange: (event: React.ChangeEvent<HTMLSelectElement>) => void
}

const KindDropdown: React.VFC<Props> = (props) => {
  return (
    <Select
      labelId="demo-simple-select-label"
      id="demo-simple-select"
      value={props.kind}
      onChange={props.handleKindChange}
    >
      <MenuItem value="random">ランダムなパスワード</MenuItem>
      <MenuItem value="simple">覚えやすいパスワード</MenuItem>
      <MenuItem value="pin">暗証番号</MenuItem>
    </Select>
  )
}

export default KindDropdown
