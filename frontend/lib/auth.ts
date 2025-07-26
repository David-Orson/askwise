import GoogleProvider from "next-auth/providers/google";
import { NextAuthOptions } from "next-auth";

export type Session = {
  id: string;
  name: string;
  email: string;
  image: string | null;
};

export const authOptions: NextAuthOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    }),
  ],
  pages: {
    signIn: "/auth/signin",
    signOut: "/auth/signout",
    error: "/auth/error",
    newUser: "/auth/new-user",
  },
  secret: process.env.NEXTAUTH_SECRET || "",
  session: {
    strategy: "jwt",
  },
  callbacks: {
    async session({ session, token }) {
      return {
        ...session,
        user: {
          ...session.user,
          id: token.sub,
        },
      };
    },
  },
};
