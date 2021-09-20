import React, { useState } from 'react'
import { Button, makeStyles } from '@material-ui/core'

interface Props {
  password: string
}

const useStyles = makeStyles({
  button: {
    color: '#ffffff',
    backgroundColor: '#0f83ec',
  },
})

const CopyButton: React.VFC<Props> = (props) => {
  const classes = useStyles()
  const [copied, setCopied] = useState(false)

  const copyPassword = () => {
    navigator.clipboard.writeText(props.password).then(() => {
      setCopied(true)
    })
  }

  // リセットのためだけど、無駄に呼ばれることになっているかも。
  setInterval(() => {
    if (copied) {
      setCopied(false)
    }
  }, 3000)

  return (
    <Button
      color="primary"
      variant="outlined"
      onClick={copyPassword}
      size="large"
      className={classes.button}
    >
      {copied ? 'パスワードをコピーしました' : '安全なパスワードをコピー'}
    </Button>
  )
}

export default CopyButton
