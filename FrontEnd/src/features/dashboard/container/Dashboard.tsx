"use client";

import DashboardBanner from "@/features/dashboard/components/DashboardBanner";
import { useDashboardCohorts } from "@/features/dashboard/hooks/useDashboard";
import { useState } from "react";
import { Cohort } from "@/features/dashboard/types/type";
import DashboardHeader from "@/features/dashboard/components/DashboardHeader";
import CohortToolbar from "@/features/dashboard/components/CohortToolbar";
import CohortCard from "@/features/dashboard/components/CohortCard";
import { SearchX } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import CreateCohortDialog from "@/features/dashboard/components/CreateCohortDialog";
import CohortPreviewDialog from "@/features/dashboard/components/CohortPreviewDialog";

const Dashboard = () => {
  const {
    filteredCohorts,
    search,
    setSearch,
    filter,
    setFilter,
    createCohort,
    resetFilters,
  } = useDashboardCohorts();

  const [createOpen, setCreateOpen] = useState(false);
  const [selectedCohort, setSelectedCohort] = useState<Cohort | null>(null);

  return (
    <div className="flex w-full min-w-0 flex-col gap-9 px-5 py-7 sm:px-8 lg:px-10">
      <DashboardBanner onCreate={() => setCreateOpen(true)} />

      <section id="classes" className="flex scroll-mt-24 flex-col gap-6">
        <DashboardHeader />

        <CohortToolbar
          filter={filter}
          onFilterChange={setFilter}
          search={search}
          onSearchChange={setSearch}
        />

        {filteredCohorts.length > 0 ? (
          <div className="grid grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-3">
            {filteredCohorts.map((cohort) => (
              <CohortCard
                key={cohort.id}
                cohort={cohort}
                onOpen={setSelectedCohort}
              />
            ))}
          </div>
        ) : (
          <div className="flex min-h-64 flex-col items-center justify-center gap-3 rounded-2xl border border-dashed border-[#DCE6ED] bg-white p-8 text-center">
            <div className="bg-secondary-light text-primary flex size-14 items-center justify-center rounded-2xl">
              <SearchX className="size-6" />
            </div>
            <h3 className="text-lg-bold text-primary-text">
              Kelas Tidak Ditemukan
            </h3>
            <p className="text-sm-medium text-secondary-text max-w-md">
              Tidak ada kelas yang cocok dengan pencarian atau filter yang
              dipilih.
            </p>

            <Button
              type="button"
              variant="outline"
              onClick={resetFilters}
              className="mt-2 cursor-pointer rounded-xl"
            >
              Reset Filter
            </Button>
          </div>
        )}
      </section>

      <CreateCohortDialog
        open={createOpen}
        onClose={() => setCreateOpen(false)}
        onCreate={createCohort}
      />

      <CohortPreviewDialog
        cohort={selectedCohort}
        onClose={() => setSelectedCohort(null)}
      />
    </div>
  );
};

export default Dashboard;
