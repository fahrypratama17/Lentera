import { BookOpen } from "lucide-react";

import MaterialItem from "./MaterialItem";
import type { MaterialWithProgress } from "../types/type";

const MaterialList = ({ materials }: { materials: MaterialWithProgress[] }) => {
  if (materials.length === 0) {
    return (
      <div className="flex min-h-48 flex-col items-center justify-center gap-3 rounded-2xl border border-dashed border-[#DCE6ED] bg-white p-8 text-center">
        <div className="bg-secondary-light text-primary flex size-12 items-center justify-center rounded-2xl">
          <BookOpen className="size-6" />
        </div>
        <p className="text-secondary-text text-sm">
          Belum ada materi yang dipublikasikan.
        </p>
      </div>
    );
  }

  const sorted = [...materials].sort((a, b) => a.position - b.position);

  return (
    <ol className="flex flex-col gap-3">
      {sorted.map((material) => (
        <li key={material.id}>
          <MaterialItem material={material} />
        </li>
      ))}
    </ol>
  );
};

export default MaterialList;