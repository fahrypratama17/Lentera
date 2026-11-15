"use client";

import { useEffect, useRef, type FormEvent } from "react";
import {
  CalendarDays,
  ExternalLink,
  Send,
  Sparkles,
  X,
} from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import { Field, FieldLabel } from "@/shared/components/ui/field";
import { formatDateTime, formatDueLabel } from "../utils/format";
import TaskStatusBadge from "./TaskStatusBadge";
import { REVIEW_DECISION_CONFIG, type TaskDialogProps } from "../types/type";

const inputClassName =
  "h-11 w-full rounded-xl border border-[#DCE6ED] bg-surface px-4 text-sm text-primary-text outline-none transition focus:border-secondary focus:ring-4 focus:ring-secondary/10";

const TaskDialog = ({
  task,
  open,
  submitting,
  error,
  onClose,
  onSubmit,
  onStart,
}: TaskDialogProps) => {
  const dialogRef = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;

    if (open && !dialog.open) {
      dialog.showModal();
    } else if (!open && dialog.open) {
      dialog.close();
    }
  }, [open]);

  const canSubmit =
    task?.status === "not_started" || task?.status === "in_progress";

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const form = event.currentTarget;
    const data = new FormData(form);

    onSubmit({
      content: String(data.get("content") ?? "").trim(),
      url: String(data.get("url") ?? "").trim(),
    });
  }

  return (
    <dialog
      ref={dialogRef}
      onClose={onClose}
      className="text-primary-text fixed inset-0 m-auto max-h-[90dvh] w-[calc(100%-2rem)] max-w-xl overflow-y-auto rounded-3xl border border-[#DCE6ED] bg-white p-0 shadow-2xl backdrop:bg-[#0A172B]/50"
    >
      {task && (
        <>
          <div className="flex items-start justify-between gap-4 border-b border-[#DCE6ED] px-6 py-5">
            <div className="min-w-0">
              <p className="text-secondary-text text-xs font-semibold">
                {task.cohortTitle} · Materi {task.position}
              </p>

              <h2 className="text-xl-bold text-primary-text mt-1">
                {task.materialTitle}
              </h2>

              <p className="text-secondary-text mt-1 flex items-center gap-1.5 text-xs">
                <CalendarDays className="size-3.5" />
                {formatDueLabel(task.deadline)}
              </p>
            </div>

            <Button
              type="button"
              variant="ghost"
              size="icon"
              onClick={onClose}
              aria-label="Tutup"
              className="text-secondary-text cursor-pointer"
            >
              <X className="size-5" />
            </Button>
          </div>

          <div className="flex flex-col gap-5 px-6 py-6">
            <div className="flex items-center justify-between gap-3">
              <span className="text-secondary-text text-xs font-semibold">
                Status
              </span>
              <TaskStatusBadge status={task.status} />
            </div>

            {task.reviewDecision && (
              <div className="bg-surface rounded-xl p-4 text-sm">
                <p
                  className={`font-semibold ${
                    REVIEW_DECISION_CONFIG[task.reviewDecision].className
                  }`}
                >
                  Review mentor: {REVIEW_DECISION_CONFIG[task.reviewDecision].label}
                </p>

                {task.reviewNote && (
                  <p className="text-primary-text mt-1 leading-relaxed">
                    {task.reviewNote}
                  </p>
                )}

                {task.reviewerName && task.reviewedAt && (
                  <p className="text-secondary-text mt-2 text-xs">
                    {task.reviewerName} · {formatDateTime(task.reviewedAt)}
                  </p>
                )}
              </div>
            )}

            {task.submissionContent && (
              <div className="bg-secondary-light rounded-xl p-4">
                <p className="text-primary mb-2 text-sm font-bold">
                  Hasil yang dikumpulkan
                </p>

                <p className="text-primary-text text-sm leading-relaxed whitespace-pre-line">
                  {task.submissionContent}
                </p>

                {task.submittedAt && (
                  <p className="text-secondary-text mt-2 text-xs">
                    Dikumpulkan {formatDateTime(task.submittedAt)}
                  </p>
                )}
              </div>
            )}

            {task.submissionUrl && (
              <a
                href={task.submissionUrl}
                target="_blank"
                rel="noopener noreferrer"
                className="text-primary flex w-fit items-center gap-1.5 text-sm font-semibold underline underline-offset-4"
              >
                <ExternalLink className="size-4" />
                Buka tautan pekerjaan
              </a>
            )}

            {error && (
              <p role="alert" className="text-destructive text-sm font-medium">
                {error}
              </p>
            )}

            {task.status === "not_started" && (
              <Button
                type="button"
                onClick={onStart}
                disabled={submitting}
                className="bg-primary hover:bg-secondary h-11 w-full cursor-pointer rounded-xl font-bold text-white"
              >
                <Sparkles className="size-4" />
                {submitting ? "Memproses..." : "Mulai kerjakan"}
              </Button>
            )}

            {canSubmit && (
              <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                <Field className="gap-2">
                  <FieldLabel htmlFor="task-content">Hasil belajar</FieldLabel>

                  <textarea
                    id="task-content"
                    name="content"
                    rows={4}
                    defaultValue={task.submissionContent ?? ""}
                    placeholder="Tulis ringkasan hasil pekerjaanmu..."
                    className={`${inputClassName} h-auto resize-y py-3`}
                  />
                </Field>

                <Field className="gap-2">
                  <FieldLabel htmlFor="task-url">
                    Tautan pekerjaan (opsional)
                  </FieldLabel>

                  <input
                    id="task-url"
                    name="url"
                    type="url"
                    defaultValue={task.submissionUrl ?? ""}
                    placeholder="https://github.com/username/repo"
                    className={inputClassName}
                  />
                </Field>

                <Button
                  type="submit"
                  disabled={submitting}
                  className="bg-primary hover:bg-secondary h-11 w-full cursor-pointer rounded-xl font-bold text-white"
                >
                  <Send className="size-4" />
                  {submitting ? "Mengumpulkan..." : "Kumpulkan hasil"}
                </Button>
              </form>
            )}

            {task.status === "submitted" && (
              <p className="text-secondary-text text-sm leading-relaxed">
                Hasilmu sedang menunggu review mentor. Pengumpulan belum bisa
                diubah sampai keputusan diberikan.
              </p>
            )}

            {task.status === "completed" && (
              <p className="text-sm leading-relaxed text-emerald-700">
                Tugas ini sudah disetujui mentor.
                {task.completedAt &&
                  ` Selesai ${formatDateTime(task.completedAt)}.`}
              </p>
            )}
          </div>
        </>
      )}
    </dialog>
  );
};

export default TaskDialog;
