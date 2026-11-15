import { CohortEmptyStateProps } from "@/features/dashboard/types/type";
import { SearchX } from "lucide-react";
import { Button } from "@/shared/components/ui/button";

const CohortEmptyState = ({ onReset }: CohortEmptyStateProps) => {
  return (
    <div className="flex min-h-64 flex-col items-center justify-center gap-3 rounded-2xl border border-dashed border-[#DCE6ED] bg-white p-8 text-center">
      <div className="bg-secondary-light text-primary flex size-14 items-center justify-center rounded-2xl">
        <SearchX className="size-6" />
      </div>
      <h3 className="text-lg-bold text-primary-text">Kelas Tidak Ditemukan</h3>
      <p className="text-sm-medium text-secondary-text max-w-md">
        Tidak ada kelas yang cocok dengan pencarian atau filter yang dipilih.
      </p>
      <Button
        type="button"
        variant="outline"
        onClick={onReset}
        className="mt-2 cursor-pointer rounded-xl"
      >
        Reset Filter
      </Button>
    </div>
  );
};

export default CohortEmptyState;
