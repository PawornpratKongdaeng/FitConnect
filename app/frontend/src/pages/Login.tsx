
import { useState } from "react";
import { api } from "../lib/api";
import { Button, Card, CardBody, Input } from "@heroui/react";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const nav = useNavigate();

  const submit = async () => {
    const { data } = await api.post("/auth/login", { email, password });
    localStorage.setItem("token", data.accessToken);
    nav("/");
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-100 via-white to-pink-100 p-4">
      <Card className="w-full max-w-md shadow-xl rounded-2xl border-0">
        <CardBody className="space-y-6 py-8 px-6">
          <div className="flex flex-col items-center gap-2">
            <div className="bg-blue-500 rounded-full p-3 mb-2 shadow-md">
              {/* Logo or icon placeholder */}
              <svg width="32" height="32" fill="none" viewBox="0 0 24 24" stroke="white"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" /></svg>
            </div>
            <h1 className="text-2xl font-bold text-gray-800">Sign in to FitConnect</h1>
            <p className="text-gray-500 text-sm">Welcome back! Please enter your details.</p>
          </div>
          <div className="space-y-4">
            <Input
              label="Email"
              value={email}
              onValueChange={setEmail}
              type="email"
              className="bg-gray-50"
            />
            <Input
              label="Password"
              value={password}
              onValueChange={setPassword}
              type="password"
              className="bg-gray-50"
            />
          </div>
          <div className="flex justify-between items-center text-sm">
            <div />
            <a href="#" className="text-blue-500 hover:underline">Forgot password?</a>
          </div>
          <Button onPress={submit} className="w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 rounded-lg shadow">Continue</Button>
        </CardBody>
      </Card>
    </div>
  );
}
