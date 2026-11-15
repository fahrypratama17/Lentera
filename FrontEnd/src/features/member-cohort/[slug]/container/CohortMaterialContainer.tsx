"use client";

import MaterialList from "../components/MaterialList";
import { useMaterialsQuery } from "@/shared/repository/materials/query";

const CohortMaterialContainer = ({ cohortId }: { cohortId: string }) => {
  const { data, isLoading, isError, refetch } = useMaterialsQuery(cohortId);

  if (isLoading) {
    return (
      <div aria-busy="true" className="flex flex-col gap-3">
        {Array.from({ length: 4 }).map((_, index) => (
          <div
            key={index}
            className="h-24 animate-pulse rounded-2xl border border-[#DCE6ED] bg-white"
          />
        ))}
      </div>
    );
  }

  if (isError) {
    return (
      <div role="alert" className="text-sm text-red-600">
        Gagal memuat materi.{" "}
        <button type="button" onClick={() => refetch()} className="underline">
          Coba lagi
        </button>
      </div>
    );
  }

  return <MaterialList materials={data ?? []} />;
};

export default CohortMaterialContainer;