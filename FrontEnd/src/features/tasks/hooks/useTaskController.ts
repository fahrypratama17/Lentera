"use client";

import { useCallback, useMemo, useState } from "react";

import { startTask, submitTask } from "../services/taskActions";
import {
  matchesFilter,
  matchesSearch,
  sortTasks,
  summarize,
} from "../utils/task";
import type {
  SubmitTaskValues,
  TaskCohortOption,
  TaskFilter,
  TaskItem,
} from "../types/type";

export const TASKS_PER_PAGE = 6;

/**
 * Satu-satunya tempat state & orkestrasi untuk halaman Tugas.
 * Komponen presentasional menerima hasilnya lewat props dari container.
 */
export function useTaskController(initialTasks: TaskItem[]) {
  const [tasks, setTasks] = useState<TaskItem[]>(() => sortTasks(initialTasks));

  const [filter, setFilterValue] = useState<TaskFilter>("all");
  const [cohortId, setCohortIdValue] = useState("all");
  const [search, setSearchValue] = useState("");
  const [page, setPage] = useState(1);

  const [selectedId, setSelectedId] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const selectedTask = useMemo(
    () => tasks.find((task) => task.id === selectedId) ?? null,
    [tasks, selectedId],
  );

  const cohortOptions = useMemo<TaskCohortOption[]>(() => {
    const unique = new Map<string, string>();
    tasks.forEach((task) => unique.set(task.cohortId, task.cohortTitle));

    return [
      { label: "Semua kelas", value: "all" },
      ...[...unique].map(([value, label]) => ({ label, value })),
    ];
  }, [tasks]);

  const filteredTasks = useMemo(
    () =>
      tasks.filter(
        (task) =>
          matchesFilter(task, filter) &&
          (cohortId === "all" || task.cohortId === cohortId) &&
          matchesSearch(task, search),
      ),
    [tasks, filter, cohortId, search],
  );

  const summary = useMemo(() => summarize(tasks), [tasks]);

  const totalItems = filteredTasks.length;
  const totalPages = Math.max(1, Math.ceil(totalItems / TASKS_PER_PAGE));
  const currentPage = Math.min(page, totalPages);

  const paginatedTasks = useMemo(() => {
    const start = (currentPage - 1) * TASKS_PER_PAGE;
    return filteredTasks.slice(start, start + TASKS_PER_PAGE);
  }, [filteredTasks, currentPage]);

  const setFilter = useCallback((value: TaskFilter) => {
    setFilterValue(value);
    setPage(1);
  }, []);

  const setCohortId = useCallback((value: string) => {
    setCohortIdValue(value);
    setPage(1);
  }, []);

  const setSearch = useCallback((value: string) => {
    setSearchValue(value);
    setPage(1);
  }, []);

  const resetFilters = useCallback(() => {
    setFilterValue("all");
    setCohortIdValue("all");
    setSearchValue("");
    setPage(1);
  }, []);

  const openTask = useCallback((task: TaskItem) => {
    setSelectedId(task.id);
    setError(null);
  }, []);

  const closeDialog = useCallback(() => {
    setSelectedId(null);
    setError(null);
  }, []);

  const replaceTask = useCallback((next: TaskItem) => {
    setTasks((previous) =>
      sortTasks(
        previous.map((task) => (task.id === next.id ? next : task)),
      ),
    );
  }, []);

  const handleStart = useCallback(async () => {
    if (!selectedTask) return;

    setSubmitting(true);
    setError(null);

    try {
      replaceTask(await startTask(selectedTask));
    } catch (cause) {
      setError(
        cause instanceof Error ? cause.message : "Gagal memulai tugas.",
      );
    } finally {
      setSubmitting(false);
    }
  }, [selectedTask, replaceTask]);

  const handleSubmit = useCallback(
    async (values: SubmitTaskValues) => {
      if (!selectedTask) return;

      setSubmitting(true);
      setError(null);

      try {
        replaceTask(await submitTask(selectedTask, values));
        setSelectedId(null);
      } catch (cause) {
        setError(
          cause instanceof Error ? cause.message : "Gagal mengumpulkan hasil.",
        );
      } finally {
        setSubmitting(false);
      }
    },
    [selectedTask, replaceTask],
  );

  return {
    /* data */
    filteredTasks,
    paginatedTasks,
    summary,
    cohortOptions,
    /* filter */
    filter,
    cohortId,
    search,
    page: currentPage,
    totalPages,
    totalItems,
    /* dialog */
    selectedTask,
    dialogOpen: selectedTask !== null,
    submitting,
    error,
    /* aksi */
    setFilter,
    setCohortId,
    setSearch,
    setPage,
    resetFilters,
    openTask,
    closeDialog,
    handleStart,
    handleSubmit,
  };
}
