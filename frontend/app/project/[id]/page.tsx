import { notFound } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { ArrowLeft } from "lucide-react";

interface ProjectPageProps {
  params: Promise<{ id: string }>;
}

export default async function ProjectPage({ params }: ProjectPageProps) {
  const { id: projectId } = await params;

  const project = {
    id: projectId,
    name: projectId === "1" ? "AI Summit" : "Tax Law Research",
    documents: [
      { id: "doc1", name: "whitepaper.pdf", updatedAt: "2 days ago" },
      { id: "doc2", name: "briefing.pdf", updatedAt: "5 days ago" },
    ],
  };

  if (!project) notFound();

  return (
    <main className="min-h-screen px-6 py-12 flex flex-col gap-8">
      {/* Back Link */}
      <Link
        href="/dashboard"
        className="flex items-center text-sm text-muted-foreground hover:text-foreground transition mb-2 w-fit"
      >
        <ArrowLeft className="w-4 h-4 mr-1" />
        Back to Dashboard
      </Link>

      <header>
        <h1 className="text-3xl font-bold mb-2">{project.name}</h1>
        <Button className="mt-2">Upload Document</Button>
      </header>

      {/* Document List */}
      <section className="grid gap-4">
        {project.documents.map((doc) => (
          <div
            key={doc.id}
            className="border rounded-lg p-4 flex justify-between items-center"
          >
            <div>
              <p className="font-medium">{doc.name}</p>
              <p className="text-sm text-muted-foreground">
                Last updated {doc.updatedAt}
              </p>
            </div>
            <Button size="sm" variant="destructive">
              Delete
            </Button>
          </div>
        ))}
      </section>

      {/* Ask Questions Section */}
      <section className="mt-8 border-t pt-6">
        <h2 className="text-xl font-semibold mb-4">Ask a Question</h2>
        <div className="flex gap-2">
          <Input
            placeholder="Ask anything about this project..."
            className="flex-1"
          />
          <Button>Submit</Button>
        </div>
      </section>
    </main>
  );
}
