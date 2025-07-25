"use client";

import { signIn } from "next-auth/react";
import { Button } from "@/components/ui/button";

export default function LoginPage() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center px-4">
      <h1 className="text-3xl font-bold mb-6">Login to AskWise</h1>
      <Button
        onClick={() => signIn("google", { callbackUrl: "/dashboard" })}
        size="lg"
      >
        Sign in with Google
      </Button>
    </div>
  );
}
