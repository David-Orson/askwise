import { Button } from "@/components/ui/button";
import Link from "next/link";
import { FeatureCard } from "@/components/FeatureCard";

export default function HomePage() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-start bg-background px-4 py-24 text-center">
      {/* Hero Section */}
      <div className="text-center max-w-2xl md:py-48">
        <h1 className="text-5xl font-bold tracking-tight mb-6">AskWise</h1>
        <p className="text-xl text-muted-foreground mb-8">
          Upload any document. Ask it questions. Get answers instantly with AI.
        </p>
        <Link href="/login">
          <Button size="lg" className="text-base">
            Get Started — It’s Free
          </Button>
        </Link>
      </div>

      {/* How It Works */}
      <section className="mt-32 px-4 max-w-4xl w-full">
        <h2 className="text-3xl font-bold mb-8">How it works</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 text-left">
          <div>
            <h3 className="text-xl font-semibold mb-2">1. Sign In</h3>
            <p className="text-muted-foreground">
              Use your Google account to get started instantly.
            </p>
          </div>
          <div>
            <h3 className="text-xl font-semibold mb-2">2. Upload</h3>
            <p className="text-muted-foreground">
              Drag and drop your PDF or doc — we’ll index it.
            </p>
          </div>
          <div>
            <h3 className="text-xl font-semibold mb-2">3. Ask</h3>
            <p className="text-muted-foreground">
              Ask questions in natural language. Get answers fast.
            </p>
          </div>
        </div>
      </section>

      {/* Feature Highlights */}
      <section className="mt-32 px-4 max-w-5xl w-full">
        <h2 className="text-3xl font-bold text-center mb-10">Why AskWise?</h2>
        <div className="grid md:grid-cols-3 gap-8">
          <FeatureCard
            title="Accurate Answers"
            description="Uses AI to extract the most relevant info from your docs."
          />
          <FeatureCard
            title="No Setup Needed"
            description="Just upload and ask — no training or tech skills required."
          />
          <FeatureCard
            title="Private & Secure"
            description="Your data stays private — always encrypted and isolated."
          />
        </div>
      </section>

      {/* Testimonials */}
      <section className="mt-32 px-4 max-w-3xl w-full">
        <h2 className="text-3xl font-bold text-center mb-10">
          What our users say
        </h2>
        <div className="space-y-8 text-left">
          <blockquote className="border-l-4 pl-4 italic text-muted-foreground">
            “This changed how I read whitepapers.” – 🧠 Fake User
          </blockquote>
          <blockquote className="border-l-4 pl-4 italic text-muted-foreground">
            “I got answers to a 200-page legal PDF in 30 seconds. Unreal.” – 🧑‍⚖️
            Definitely Real Lawyer
          </blockquote>
          <blockquote className="border-l-4 pl-4 italic text-muted-foreground">
            “I use AskWise to summarize grant docs while I drink coffee.” – ☕
            Funded Founder
          </blockquote>
        </div>
      </section>

      {/* Footer */}
      <footer className="mt-32 text-center text-sm text-muted-foreground border-t pt-8 w-full max-w-2xl">
        <p>© {new Date().getFullYear()} AskWise</p>
        <div className="flex justify-center gap-6 mt-2">
          <Link href="/privacy" className="hover:underline">
            Privacy
          </Link>
          <Link href="/terms" className="hover:underline">
            Terms
          </Link>
          <a
            href="https://github.com/David-Orson/askwise"
            target="_blank"
            rel="noopener noreferrer"
            className="hover:underline"
          >
            GitHub
          </a>
        </div>
      </footer>
    </main>
  );
}
