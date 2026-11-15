import CohortCard from "@/shared/components/Cohort/CohortCard";
import type { CohortGridProps } from "../types/cohort-type";

const CohortGrid = ({ cohorts, onOpen }: CohortGridProps) => {
  return (
    <div className="grid grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-3">
      {cohorts.map((cohort) => (
        <CohortCard key={cohort.id} cohort={cohort} onOpen={onOpen} />
      ))}
    </div>
  );
};

export default CohortGrid;