"use client";

import CohortActionBar from "../components/CohortActionBar";
import CohortGrid from "../components/CohortGrid";
import CohortPagination from "../components/CohortPagination";
import CohortEmptyState from "@/shared/components/CohortEmptyState";
import CreateCohortDialog from "@/shared/components/CreateCohortDialog";
import { useCohortController } from "../hooks/useCohort";
import { COHORTS_PER_PAGE } from "../hooks/useCohort";
import CohortBannerHeader from "../components/CohortBannerHeader";

const CohortContainer = () => {
  const {
    paginatedCohorts,
    totalItems,
    totalPages,
    currentPage,
    setPage,
    search,
    setSearch,
    filter,
    setFilter,
    resetFilters,
    createCohort,
    createOpen,
    openCreateDialog,
    closeCreateDialog,
    openCohort,
  } = useCohortController();

  return (
    <div className="flex w-full min-w-0 flex-col gap-7 px-5 py-7 sm:px-8 lg:px-10">
      <CohortBannerHeader />

      <CohortActionBar
        filter={filter}
        onFilterChange={setFilter}
        search={search}
        onSearchChange={setSearch}
        onCreate={openCreateDialog}
      />

      {totalItems > 0 ? (
        <>
          <CohortGrid cohorts={paginatedCohorts} onOpen={openCohort} />

          <CohortPagination
            currentPage={currentPage}
            totalPages={totalPages}
            totalItems={totalItems}
            pageSize={COHORTS_PER_PAGE}
            onPageChange={setPage}
          />
        </>
      ) : (
        <CohortEmptyState onReset={resetFilters} />
      )}

      <CreateCohortDialog
        open={createOpen}
        onClose={closeCreateDialog}
        onCreate={createCohort}
      />
    </div>
  );
};

export default CohortContainer;
