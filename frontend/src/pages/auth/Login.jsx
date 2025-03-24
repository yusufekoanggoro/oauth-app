import { useState } from "react";
import LoginForm from "../../components/LoginForm";
import GoogleLoginButton from "../../components/GoogleLoginButton";
// import { googleLogout, useGoogleLogin } from '@react-oauth/google';

export default function Login() {
  const [isLoggingIn, setIsLoggingIn] = useState(false);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 shadow-lg rounded-xl w-96 flex flex-col items-center"> 
        <h2 className="text-2xl font-bold text-center mb-4">
          Masuk ke Akun
        </h2>

        {!isLoggingIn ? (
          <>
            <GoogleLoginButton buttonText="Masuk dengan Google" />
          </>
        ) : (
          <LoginForm />
        )}
      </div>
    </div>
  );
}
