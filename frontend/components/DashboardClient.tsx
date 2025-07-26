"use client";

import { useSession } from "next-auth/react";
import { signOut } from "next-auth/react";
import { useEffect, useRef } from "react";
import { useSyncUser } from "@/hooks/useSyncUser";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function DashboardClient({
  serverSession,
}: {
  serverSession: any;
}) {
  const { status } = useSession();
  const syncUser = useSyncUser();
  const hasSyncedRef = useRef(false);

  useEffect(() => {
    if (
      status === "authenticated" &&
      !hasSyncedRef.current &&
      !syncUser.isPending &&
      !syncUser.isSuccess &&
      !syncUser.isError
    ) {
      hasSyncedRef.current = true;
      syncUser.mutate(undefined, {
        onError: (error) => {
          console.error("User sync failed:", error);
          // Optional: show UI message or toast here
          signOut({ callbackUrl: "/" }); // 👈 log out and redirect
        },
      });
    }
  }, [status, syncUser]);

  const projects = [
    { id: "1", name: "AI Summit" },
    { id: "2", name: "Tax Law Research" },
  ];

  return (
    <main className="min-h-screen px-6 py-12">
      <h1 className="text-3xl font-bold mb-4">
        Welcome, {serverSession.user?.name?.split(" ")[0]} 👋
      </h1>
      <p className="text-muted-foreground mb-6">
        Organize your documents into projects and start asking questions.
      </p>

      <Button className="mb-8">New Project</Button>

      <div className="grid gap-4">
        {projects.map((project) => (
          <Link
            key={project.id}
            href={`/project/${project.id}`}
            className="border rounded-lg p-4 flex justify-between items-center hover:bg-muted transition"
          >
            <span className="font-medium">{project.name}</span>
            <span className="text-sm text-muted-foreground">View →</span>
          </Link>
        ))}
      </div>
    </main>
  );
}
