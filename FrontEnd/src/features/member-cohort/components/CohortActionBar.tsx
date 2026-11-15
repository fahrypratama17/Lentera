import { Plus } from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import CohortToolbar from "@/shared/components/CohortToolbar";
import type { CohortActionBarProps } from "../types/cohort-type";

const CohortActionBar = ({
  filter,
  onFilterChange,
  search,
  onSearchChange,
  onCreate,
}: CohortActionBarProps) => {
  return (
    <div className="flex flex-col gap-4 lg:flex-row lg:items-center">
      <div className="min-w-0 flex-1">
        <CohortToolbar
          filter={filter}
          onFilterChange={onFilterChange}
          search={search}
          onSearchChange={onSearchChange}
        />
      </div>

      <Button
        type="button"
        onClick={onCreate}
        className="bg-primary hover:bg-secondary h-11 shrink-0 cursor-pointer rounded-xl px-5 font-bold text-white"
      >
        <Plus className="size-4" />
        Buat Kelas
      </Button>
    </div>
  );
};

export default CohortActionBar;
