import { useQuery } from "@tanstack/react-query";
import { api } from "../lib/api";
import { Card, CardBody, Chip } from "@heroui/react";

type Trainer = {
  id: string;
  bio: string;
  specialties: string[];
  rate: number;
  regionCode: string;
};

export default function Trainers() {
  const { data } = useQuery({
    queryKey: ["trainers", "bkk"],
    queryFn: async () => {
      const res = await api.get("/trainers", {
        params: { region: "BKK", limit: 20 },
      });
      return res.data as Trainer[];
    },
  });

  return (
    <div className="max-w-5xl mx-auto p-4 grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
      {data?.map((t) => (
        <Card key={t.id}>
          <CardBody className="space-y-2">
            <div className="font-semibold">{t.bio}</div>
            <div className="flex gap-2 flex-wrap">
              {t.specialties?.map((s) => (
                <Chip key={s}>{s}</Chip>
              ))}
            </div>
            <div className="text-sm text-neutral-500">
              {t.regionCode} • ฿{t.rate}/hr
            </div>
          </CardBody>
        </Card>
      ))}
    </div>
  );
}
