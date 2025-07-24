import "@/styles/globals.css";
import { ReactQueryProvider } from "@/lib/react-query-provider";

export const metadata = {
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
        <ReactQueryProvider>{children}</ReactQueryProvider>
      </body>
    </html>
  );
}
