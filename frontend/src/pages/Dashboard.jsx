import { useNavigate } from "react-router-dom";
import { googleLogout } from "@react-oauth/google";

export default function Dashboard() {
  const navigate = useNavigate();

  const handleLogout = () => {
    googleLogout();
    navigate("/");
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
      <div className="bg-white shadow-lg rounded-lg p-6 w-full max-w-md text-center">
        <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
        <p className="text-gray-600 mb-4">Selamat datang di dashboard!</p>
        <button onClick={handleLogout}>Log out</button>
      </div>
    </div>
  );
}
