import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import { useRouter } from "next/router";
import { auth, firebaseUser, uiConfig } from "../lib/firebase";

const sign_in: React.VFC = () => {
  const router = useRouter();
  if (firebaseUser()) {
    router.push("/");
  }

  return <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />;
};

export default sign_in;
