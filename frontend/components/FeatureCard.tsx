interface Props {
  title: string;
  description: string;
}

export function FeatureCard({ title, description }: Props) {
  return (
    <div className="border rounded-xl p-6 text-left shadow-sm bg-white dark:bg-zinc-900">
      <h3 className="text-lg font-semibold mb-2">{title}</h3>
      <p className="text-muted-foreground">{description}</p>
    </div>
  );
}
