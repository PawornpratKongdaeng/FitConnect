import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { api } from "../lib/api";

export default function GymsNear() {
  const [coords, setCoords] = useState<{ lat: number; lng: number }>();

  useEffect(() => {
    navigator.geolocation.getCurrentPosition((p) =>
      setCoords({ lat: p.coords.latitude, lng: p.coords.longitude })
    );
  }, []);

  const { data } = useQuery({
    enabled: !!coords,
    queryKey: ["gyms", coords?.lat, coords?.lng],
    queryFn: async () => {
      const res = await api.get("/gyms/near", {
        params: { lat: coords!.lat, lng: coords!.lng, km: 5 },
      });
      return res.data as { id: string; name: string; province: string }[];
    },
  });

  return (
    <div className="max-w-3xl mx-auto p-4 space-y-3">
      <h1 className="text-xl font-semibold">Gyms Near Me</h1>
      <ul className="space-y-2">
        {data?.map((g) => (
          <li key={g.id} className="p-3 rounded-xl border">
            {g.name} • {g.province}
          </li>
        ))}
      </ul>
    </div>
  );
}
