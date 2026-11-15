"use client";

import TaskBanner from "../components/TaskBanner";
import TaskSummary from "../components/TaskSummary";
import TaskToolbar from "../components/TaskToolbar";
import TaskList from "../components/TaskList";
import TaskPagination from "../components/TaskPagination";
import TaskEmptyState from "../components/TaskEmptyState";
import TaskDialog from "../components/TaskDialog";
import { TASKS_PER_PAGE, useTaskController } from "../hooks/useTaskController";
import type { TaskItem } from "../types/type";

const TaskContainer = ({ initialTasks }: { initialTasks: TaskItem[] }) => {
  const {
    paginatedTasks,
    summary,
    cohortOptions,
    filter,
    cohortId,
    search,
    page,
    totalPages,
    totalItems,
    selectedTask,
    dialogOpen,
    submitting,
    error,
    setFilter,
    setCohortId,
    setSearch,
    setPage,
    resetFilters,
    openTask,
    closeDialog,
    handleStart,
    handleSubmit,
  } = useTaskController(initialTasks);

  return (
    <div className="flex w-full min-w-0 flex-col gap-7 px-5 py-7 sm:px-8 lg:px-10">
      <TaskBanner />

      <TaskSummary summary={summary} />

      <TaskToolbar
        filter={filter}
        cohortId={cohortId}
        search={search}
        cohortOptions={cohortOptions}
        onFilterChange={setFilter}
        onCohortChange={setCohortId}
        onSearchChange={setSearch}
      />

      {totalItems > 0 ? (
        <>
          <TaskList tasks={paginatedTasks} onOpen={openTask} />

          <TaskPagination
            currentPage={page}
            totalPages={totalPages}
            totalItems={totalItems}
            pageSize={TASKS_PER_PAGE}
            onPageChange={setPage}
          />
        </>
      ) : (
        <TaskEmptyState onReset={resetFilters} />
      )}

      <TaskDialog
        task={selectedTask}
        open={dialogOpen}
        submitting={submitting}
        error={error}
        onClose={closeDialog}
        onSubmit={handleSubmit}
        onStart={handleStart}
      />
    </div>
  );
};

export default TaskContainer;
