"use client";

import { useEffect, useMemo, useState } from "react";
import { mockCohorts } from "@/features/dashboard/data/data";
import type {
  Cohort,
  CohortFilter,
  CohortTone,
  CreateCohortValues,
} from "@/features/dashboard/types/type";

const STORAGE_KEY = "studytrack-cohorts-demo";

const tones: CohortTone[] = ["ocean", "teal", "slate"];

export function useDashboardCohorts() {
  const [cohorts, setCohorts] = useState<Cohort[]>(mockCohorts);

  const [search, setSearch] = useState("");
  const [filter, setFilter] = useState<CohortFilter>("all");

  const [hydrated, setHydrated] = useState(false);

  useEffect(() => {
    try {
      const stored = localStorage.getItem(STORAGE_KEY);

      if (stored) {
        const parsed: unknown = JSON.parse(stored);

        if (Array.isArray(parsed)) {
          setCohorts(parsed as Cohort[]);
        }
      }
    } catch (error) {
      console.error("Failed to load cohorts:", error);
    } finally {
      setHydrated(true);
    }
  }, []);

  useEffect(() => {
    if (!hydrated) return;

    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(cohorts));
    } catch (error) {
      console.error("Failed to save cohorts:", error);
    }
  }, [cohorts, hydrated]);

  const filteredCohorts = useMemo(() => {
    const keyword = search.trim().toLowerCase();

    return cohorts.filter((cohort) => {
      const matchesRole = filter === "all" || cohort.role === filter;

      const searchable = [
        cohort.title,
        cohort.mentorName,
        cohort.description,
        cohort.previewQuestion,
      ]
        .join(" ")
        .toLowerCase();

      return matchesRole && searchable.includes(keyword);
    });
  }, [cohorts, search, filter]);

  const counts = useMemo(
    () => ({
      all: cohorts.length,
      mentor: cohorts.filter((cohort) => cohort.role === "mentor").length,
      member: cohorts.filter((cohort) => cohort.role === "member").length,
    }),
    [cohorts],
  );

  function createCohort(values: CreateCohortValues) {
    const newCohort: Cohort = {
      id: crypto.randomUUID(),
      title: values.title.trim(),
      description: values.description.trim(),
      mentorName: values.mentorName.trim(),
      period: values.period.trim(),
      previewQuestion: values.previewQuestion.trim(),
      role: "mentor",
      memberCount: 1,
      materialCount: values.previewQuestion.trim() ? 1 : 0,
      tone: tones[cohorts.length % tones.length],
    };

    setCohorts((previous) => [newCohort, ...previous]);

    setFilter("all");
    setSearch("");
  }

  function resetFilters() {
    setFilter("all");
    setSearch("");
  }

  return {
    cohorts,
    filteredCohorts,
    counts,
    search,
    setSearch,
    filter,
    setFilter,
    createCohort,
    resetFilters,
  };
}
