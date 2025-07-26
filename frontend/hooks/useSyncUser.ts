import { useSession } from "next-auth/react";
import { useMutation } from "@tanstack/react-query";
import { Session } from "@/lib/auth";

export function useSyncUser() {
  const { data: session } = useSession();

  return useMutation({
    mutationFn: async () => {
      const user = session?.user as Session;
      if (!user.id) throw new Error("No session user ID");

      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/sync`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${user?.id}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: user.name,
          email: user.email,
          image: user.image,
        }),
      });

      if (!res.ok) throw new Error("User sync failed");
      return res.json();
    },
  });
}
