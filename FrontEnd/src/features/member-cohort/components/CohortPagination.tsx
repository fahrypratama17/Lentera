import { ChevronLeft, ChevronRight } from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import type { CohortPaginationProps } from "../types/cohort-type";

const getPageNumbers = (current: number, total: number) => {
  if (total <= 5) return Array.from({ length: total }, (_, i) => i + 1);

  const pages: (number | "...")[] = [1];
  const from = Math.max(2, current - 1);
  const to = Math.min(total - 1, current + 1);
  if (from > 2) pages.push("...");
  for (let p = from; p <= to; p++) pages.push(p);
  if (to < total - 1) pages.push("...");

  pages.push(total);
  return pages;
};

const CohortPagination = ({
  currentPage,
  totalPages,
  totalItems,
  pageSize,
  onPageChange,
}: CohortPaginationProps) => {
  if (totalPages <= 1) return null;

  const from = (currentPage - 1) * pageSize + 1;
  const to = Math.min(currentPage * pageSize, totalItems);

  return (
    <nav
      aria-label="Navigasi halaman kelas"
      className="flex flex-col items-center justify-between gap-4 sm:flex-row my-2"
    >
      <p className="text-secondary-text text-sm">
        Menampilkan {from}–{to} dari {totalItems} kelas
      </p>

      <div className="flex items-center gap-1">
        <Button
          type="button"
          variant="outline"
          size="icon"
          aria-label="Halaman sebelumnya"
          disabled={currentPage === 1}
          onClick={() => onPageChange(currentPage - 1)}
          className="size-9 cursor-pointer rounded-lg border-[#DCE6ED]"
        >
          <ChevronLeft className="size-4" />
        </Button>

        {getPageNumbers(currentPage, totalPages).map((item, index) =>
          item === "..." ? (
            <span
              key={`ellipsis-${index}`}
              className="text-secondary-text w-6 text-center text-sm"
            >
              …
            </span>
          ) : (
            <Button
              key={item}
              type="button"
              variant="ghost"
              size="sm"
              aria-current={item === currentPage ? "page" : undefined}
              onClick={() => onPageChange(item)}
              className={`size-9 cursor-pointer rounded-lg text-sm font-semibold ${
                item === currentPage
                  ? "bg-primary hover:bg-secondary text-white hover:text-white"
                  : "text-secondary-text hover:bg-secondary-light hover:text-primary"
              }`}
            >
              {item}
            </Button>
          ),
        )}

        <Button
          type="button"
          variant="outline"
          size="icon"
          aria-label="Halaman berikutnya"
          disabled={currentPage === totalPages}
          onClick={() => onPageChange(currentPage + 1)}
          className="size-9 cursor-pointer rounded-lg border-[#DCE6ED]"
        >
          <ChevronRight className="size-4" />
        </Button>
      </div>
    </nav>
  );
};

export default CohortPagination;