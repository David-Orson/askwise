import "@/styles/globals.css";
import { Providers } from "@/components/Providers";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "AskWise",
  description: "Document Q&A SaaS",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
