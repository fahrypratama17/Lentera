import { ChevronLeft, ChevronRight } from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import type { TaskPaginationProps } from "../types/type";

const TaskPagination = ({
  currentPage,
  totalPages,
  totalItems,
  pageSize,
  onPageChange,
}: TaskPaginationProps) => {
  if (totalPages <= 1) return null;

  const start = (currentPage - 1) * pageSize + 1;
  const end = Math.min(currentPage * pageSize, totalItems);

  return (
    <div className="flex flex-col items-center justify-between gap-3 sm:flex-row">
      <p className="text-secondary-text text-sm">
        Menampilkan {start}–{end} dari {totalItems} tugas
      </p>

      <div className="flex items-center gap-2">
        <Button
          type="button"
          variant="outline"
          size="sm"
          disabled={currentPage === 1}
          onClick={() => onPageChange(currentPage - 1)}
          className="h-9 cursor-pointer rounded-lg border-[#DCE6ED] px-3"
        >
          <ChevronLeft className="size-4" />
          Sebelumnya
        </Button>

        <span className="text-secondary-text text-sm font-semibold">
          {currentPage} / {totalPages}
        </span>

        <Button
          type="button"
          variant="outline"
          size="sm"
          disabled={currentPage === totalPages}
          onClick={() => onPageChange(currentPage + 1)}
          className="h-9 cursor-pointer rounded-lg border-[#DCE6ED] px-3"
        >
          Berikutnya
          <ChevronRight className="size-4" />
        </Button>
      </div>
    </div>
  );
};

export default TaskPagination;
