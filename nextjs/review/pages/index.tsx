import { Button, Grid } from '@mui/material'
import Link from 'next/link'

export const Home = (): JSX.Element => {
  return (
    <>
      <h1>Next sample!</h1>

      <Grid container justifyItems="center">
        <Grid item xs={4}>
          <Link href="/hello">
            <Button variant="contained">ハロー画面</Button>
          </Link>
        </Grid>

        <Grid item xs={4}>
          <Link href="/users">
            <Button variant="contained">ユーザ一覧</Button>
          </Link>
        </Grid>
      </Grid>
    </>
  )
}

export default Home
