import { Search } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { Input } from "@/shared/components/ui/input";
import { CohortToolbarProps, filters } from "@/features/dashboard/types/type";

const CohortToolbar = ({
  filter,
  onFilterChange,
  search,
  onSearchChange,
}: CohortToolbarProps) => {
  return (
    <div className="flex flex-col justify-between gap-4 sm:flex-row sm:items-center">
      <div
        role="group"
        aria-label="Filter kelas berdasarkan role"
        className="flex w-fit items-center gap-1 rounded-xl border border-[#DCE6ED] bg-white p-1"
      >
        {filters.map((item) => {
          const active = filter === item.value;

          return (
            <Button
              key={item.value}
              type="button"
              variant="ghost"
              size="sm"
              aria-pressed={active}
              onClick={() => onFilterChange(item.value)}
              className={`h-9 cursor-pointer rounded-lg px-4 text-sm font-semibold transition-colors ${
                active
                  ? "bg-primary hover:bg-secondary text-white hover:text-white"
                  : "text-secondary-text hover:bg-secondary-light hover:text-primary"
              } `}
            >
              {item.label}
            </Button>
          );
        })}
      </div>

      <div className="relative w-full sm:max-w-77.5">
        <Search className="text-secondary-text pointer-events-none absolute top-1/2 left-3 size-4 -translate-y-1/2" />

        <Input
          type="search"
          value={search}
          onChange={(event) => onSearchChange(event.target.value)}
          placeholder="Cari nama kelas atau mentor..."
          aria-label="Cari kelas"
          className="text-primary-text placeholder:text-secondary-text/70 focus-visible:border-secondary focus-visible:ring-secondary/15 h-11 w-full rounded-xl border-[#DCE6ED] bg-white pl-10 text-sm shadow-none"
        />
      </div>
    </div>
  );
};

export default CohortToolbar;
