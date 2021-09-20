import { Button } from '@mui/material'
import Link from 'next/link'

export const Home = (): JSX.Element => {
  return (
    <>
      <h1>next sample!</h1>
      <Link href="/hello">
        <Button variant="contained">このボタンを押してね！</Button>
      </Link>
    </>
  )
}

export default Home
