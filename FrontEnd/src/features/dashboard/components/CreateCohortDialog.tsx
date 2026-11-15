"use client";

import { useEffect, useRef, type FormEvent } from "react";
import { Plus, X } from "lucide-react";

import { Button } from "@/shared/components/ui/button";
import { Field, FieldLabel } from "@/shared/components/ui/field";

import type { CreateCohortValues } from "@/features/dashboard/types/type";

interface CreateCohortDialogProps {
  open: boolean;
  onClose: () => void;
  onCreate: (values: CreateCohortValues) => void;
}

const inputClassName =
  "h-11 w-full rounded-xl border border-[#DCE6ED] bg-surface px-4 text-sm text-primary-text outline-none transition focus:border-secondary focus:ring-4 focus:ring-secondary/10";

const CreateCohortDialog = ({
  open,
  onClose,
  onCreate,
}: CreateCohortDialogProps) => {
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

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const form = event.currentTarget;
    const data = new FormData(form);

    const values: CreateCohortValues = {
      title: String(data.get("title") ?? "").trim(),
      mentorName: String(data.get("mentorName") ?? "").trim(),
      description: String(data.get("description") ?? "").trim(),
      period: String(data.get("period") ?? "").trim(),
      previewQuestion: String(data.get("previewQuestion") ?? "").trim(),
    };

    if (!values.title || !values.mentorName) {
      return;
    }

    onCreate(values);
    form.reset();
    onClose();
  }

  return (
    <dialog
      ref={dialogRef}
      onClose={onClose}
      className="text-primary-text fixed inset-0 m-auto max-h-[90dvh] w-[calc(100%-2rem)] max-w-xl overflow-y-auto rounded-3xl border border-[#DCE6ED] bg-white p-0 shadow-2xl backdrop:bg-[#0A172B]/50"
    >
      <div className="flex items-center justify-between border-b border-[#DCE6ED] px-6 py-5">
        <div>
          <h2 className="text-xl-bold text-primary-text">Buat Kelas Baru</h2>

          <p className="text-secondary-text mt-1 text-sm">
            Isi informasi cohort yang ingin dibuat.
          </p>
        </div>

        <Button
          type="button"
          variant="ghost"
          size="icon"
          onClick={onClose}
          aria-label="Tutup formulir"
          className="text-secondary-text cursor-pointer"
        >
          <X className="size-5" />
        </Button>
      </div>

      <form onSubmit={handleSubmit} className="flex flex-col gap-5 px-6 py-6">
        <Field className="gap-2">
          <FieldLabel htmlFor="cohort-title">Nama Kelas</FieldLabel>

          <input
            id="cohort-title"
            name="title"
            placeholder="Contoh: Frontend Development"
            className={inputClassName}
            required
            minLength={3}
            maxLength={100}
          />
        </Field>

        <Field className="gap-2">
          <FieldLabel htmlFor="mentor-name">Nama Mentor</FieldLabel>

          <input
            id="mentor-name"
            name="mentorName"
            placeholder="Nama mentor kelas"
            className={inputClassName}
            required
            maxLength={100}
          />
        </Field>

        <Field className="gap-2">
          <FieldLabel htmlFor="cohort-period">Periode Belajar</FieldLabel>

          <input
            id="cohort-period"
            name="period"
            placeholder="Contoh: November - Desember 2026"
            className={inputClassName}
          />
        </Field>

        <Field className="gap-2">
          <FieldLabel htmlFor="cohort-description">Deskripsi Kelas</FieldLabel>

          <textarea
            id="cohort-description"
            name="description"
            rows={3}
            placeholder="Jelaskan tujuan pembelajaran..."
            className={`${inputClassName} h-auto resize-y py-3`}
          />
        </Field>

        <Field className="gap-2">
          <FieldLabel htmlFor="preview-question">
            Preview Soal / Materi
          </FieldLabel>

          <textarea
            id="preview-question"
            name="previewQuestion"
            rows={3}
            placeholder="Contoh: Buat sebuah website menggunakan Next.js..."
            className={`${inputClassName} h-auto resize-y py-3`}
          />
        </Field>

        <div className="flex justify-end gap-3 border-t border-[#DCE6ED] pt-5">
          <Button
            type="button"
            variant="outline"
            onClick={onClose}
            className="h-11 cursor-pointer rounded-xl px-5"
          >
            Batal
          </Button>

          <Button
            type="submit"
            className="bg-primary hover:bg-secondary h-11 cursor-pointer rounded-xl px-5 font-bold text-white"
          >
            <Plus className="size-4" />
            Buat Kelas
          </Button>
        </div>
      </form>
    </dialog>
  );
};

export default CreateCohortDialog;
