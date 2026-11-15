import { useEffect, useRef } from "react";
import type { CohortPreviewDialogProps } from "../types/type";

export function useCohortPreviewDialog(
  cohort: CohortPreviewDialogProps["cohort"],
) {
  const dialogRef = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;

    if (cohort && !dialog.open) {
      dialog.showModal();
    } else if (!cohort && dialog.open) {
      dialog.close();
    }
  }, [cohort]);

  const closeDialog = () => {
    dialogRef.current?.close();
  };

  return {
    dialogRef,
    closeDialog,
  };
}
