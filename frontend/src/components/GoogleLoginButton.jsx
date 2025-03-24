import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";

export default function GoogleLoginButton({ buttonText }) {
  const clientId = import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID;
  const navigate = useNavigate();

  const verifyToken = async (token) => {
    try {
      const response = await fetch("http://localhost:8080/auth/google", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({token}),
      });
      const data = await response.json();
      if (data.success) {
        navigate("/dashboard");
      } else {
        console.error("Token verification failed:", data.message);
      }
    } catch (error) {
      console.error("Error verifying token:", error);
    }
  };


  return (
    <GoogleOAuthProvider clientId={clientId}>
        <GoogleLogin
          onSuccess={(response) => {
            verifyToken(response.credential);
          }}
          onError={() => {
            console.log("Login Failed");
          }}
          theme="filled_blue"
          shape="pill"
          size="large"
          text="continue_with"
          prompt="select_account"
        >

        </GoogleLogin>
    </GoogleOAuthProvider>
  );
}
