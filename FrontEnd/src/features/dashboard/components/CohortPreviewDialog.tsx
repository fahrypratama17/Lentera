import { BookOpen, UsersRound, X } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { CohortPreviewDialogProps } from "@/features/dashboard/types/type";
import { useCohortPreviewDialog } from "@/features/dashboard/hooks/useCohortPreviewDialog";

const CohortPreviewDialog = ({ cohort, onClose }: CohortPreviewDialogProps) => {
  const { dialogRef, closeDialog } = useCohortPreviewDialog(cohort);

  return (
    <dialog
      ref={dialogRef}
      onClose={onClose}
      className="fixed inset-0 m-auto max-h-[90dvh] w-[calc(100%-2rem)] max-w-lg overflow-y-auto rounded-3xl border border-[#DCE6ED] bg-white p-0 shadow-2xl backdrop:bg-[#0A172B]/50"
    >
      {cohort && (
        <>
          <div className="bg-[linear-gradient(110deg,#1F506D,#0D6684)] p-6 text-white">
            <div className="flex items-start justify-between gap-4">
              <div className="space-y-2">
                <span className="text-xs font-semibold text-white/75">
                  {cohort.role === "mentor"
                    ? "Sebagai Mentor"
                    : "Sebagai Member"}
                </span>
                <h2 className="text-2xl-bold text-white">{cohort.title}</h2>
                <p className="text-sm text-white/85">
                  Mentor: {cohort.mentorName}
                </p>
              </div>
              <Button
                type="button"
                variant="ghost"
                size="icon"
                onClick={closeDialog}
                aria-label="Tutup"
                className="cursor-pointer text-white hover:bg-white/15 hover:text-white"
              >
                <X className="size-5" />
              </Button>
            </div>
          </div>

          <div className="space-y-5 p-6">
            <p className="text-secondary-text text-sm leading-relaxed">
              {cohort.description || "Belum ada deskripsi kelas."}
            </p>

            <div className="text-secondary-text flex flex-wrap gap-4 text-sm">
              <span className="flex items-center gap-2">
                <UsersRound className="size-4" />
                {cohort.memberCount} Anggota
              </span>
              <span className="flex items-center gap-2">
                <BookOpen className="size-4" />
                {cohort.materialCount} Materi
              </span>
            </div>

            {cohort.period && (
              <p className="text-secondary-text text-sm">
                Periode: {cohort.period}
              </p>
            )}

            <div className="bg-secondary-light rounded-xl p-4">
              <p className="text-primary mb-2 text-sm font-bold">
                Preview Materi / Tugas
              </p>
              <p className="text-primary-text text-sm leading-relaxed">
                {cohort.previewQuestion ||
                  "Belum ada materi yang dipublikasikan."}
              </p>
            </div>

            <Button
              type="button"
              onClick={onClose}
              className="bg-primary hover:bg-secondary h-11 w-full cursor-pointer rounded-xl text-white"
            >
              Tutup Ringkasan
            </Button>
          </div>
        </>
      )}
    </dialog>
  );
};

export default CohortPreviewDialog;
