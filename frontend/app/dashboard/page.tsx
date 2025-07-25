import { getServerSession } from "next-auth";
import { authOptions } from "@/lib/auth";
import { redirect } from "next/navigation";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default async function DashboardPage() {
  const session = await getServerSession(authOptions);
  if (!session) redirect("/login");

  const projects = [
    { id: "1", name: "AI Summit" },
    { id: "2", name: "Tax Law Research" },
  ];

  return (
    <main className="min-h-screen px-6 py-12">
      <h1 className="text-3xl font-bold mb-4">
        Welcome, {session.user?.name?.split(" ")[0]} 👋
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
