import { useState } from "react";
import RegisterForm from "../../components/RegisterForm";
import GoogleLoginButton from "../../components/GoogleLoginButton";

export default function Register() {
  const [isRegistering, setIsRegistering] = useState(false);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 shadow-lg rounded-xl w-96">
        <h2 className="text-2xl font-bold text-center mb-4">
          Daftar Akun
        </h2>
        
        {!isRegistering ? (
          <>
            <GoogleLoginButton buttonText="Daftar dengan Google" />
            <div className="text-center my-4 text-gray-500">atau</div>
            <button 
              onClick={() => setIsRegistering(true)}
              className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600"
            >
              Daftar dengan Email
            </button>
          </>
        ) : (
          <RegisterForm />
        )}
      </div>
    </div>
  );
}
