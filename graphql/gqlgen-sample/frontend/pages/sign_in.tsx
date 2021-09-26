import React from 'react'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import { auth, uiConfig } from '../src/lib/firebase'
import { useRouter } from 'next/router'

const sign_in: React.VFC = () => {
  const router = useRouter()
  if (auth.currentUser) {
    router.push('/')
  }

  return <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
}

export default sign_in
