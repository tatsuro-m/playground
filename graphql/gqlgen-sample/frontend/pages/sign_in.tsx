import React from 'react'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import { auth, uiConfig } from '../src/lib/firebase'

const sign_in: React.VFC = () => {
  return <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
}

export default sign_in
