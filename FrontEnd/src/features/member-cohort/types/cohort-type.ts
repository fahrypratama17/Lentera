import type {
  CohortCardProps,
  CohortToolbarProps,
} from "@/features/dashboard/types/type";

export type CohortActionBarProps = CohortToolbarProps & {
  onCreate: () => void;
};

export interface CohortGridProps {
  cohorts: CohortCardProps["cohort"][];
  onOpen: CohortCardProps["onOpen"];
}

export interface CohortPaginationProps {
  currentPage: number;
  totalPages: number;
  totalItems: number;
  pageSize: number;
  onPageChange: (page: number) => void;
}