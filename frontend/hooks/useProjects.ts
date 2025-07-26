import { useQuery } from "@tanstack/react-query";
import { useSession } from "next-auth/react";
import { Session } from "@/lib/auth";

export function useProjects() {
  const { data: session } = useSession();
  const user = session?.user as Session;

  return useQuery({
    queryKey: ["projects", user.id],
    enabled: !!user.id,
    queryFn: async () => {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/projects`, {
        headers: {
          Authorization: `Bearer ${user.id}`,
        },
      });
      if (!res.ok) throw new Error("Failed to fetch projects");
      return res.json();
    },
  });
}
