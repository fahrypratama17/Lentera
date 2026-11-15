import type { ReviewDecision, TaskStatus } from "../types/type";

export interface TaskSeed {
  cohortId: string;
  materialIndex: number;
  materialTitle: string;
  position: number;
  deadlineOffset: number | null;
  status: TaskStatus;
  submissionContent: string | null;
  submissionUrl: string | null;
  reviewDecision: ReviewDecision | null;
  reviewNote: string | null;
  reviewerName: string | null;
}

export const taskSeeds: TaskSeed[] = [
  {
    cohortId: "backend-engineering",
    materialIndex: 1,
    materialTitle: "Pengenalan & Persiapan",
    position: 1,
    deadlineOffset: -6,
    status: "completed",
    submissionContent:
      "Rangkuman arsitektur REST API dan struktur folder proyek.",
    submissionUrl: "https://github.com/bagas/backend-engineering",
    reviewDecision: "approved",
    reviewNote: "Struktur rapi, lanjut ke materi berikutnya.",
    reviewerName: "Nadia Putri",
  },
  {
    cohortId: "backend-engineering",
    materialIndex: 2,
    materialTitle: "Konsep Dasar",
    position: 2,
    deadlineOffset: -2,
    status: "in_progress",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: "revision",
    reviewNote:
      "Tambahkan penjelasan pada bagian kesimpulan, lalu kumpulkan ulang.",
    reviewerName: "Nadia Putri",
  },
  {
    cohortId: "backend-engineering",
    materialIndex: 3,
    materialTitle: "Praktik Terarah",
    position: 3,
    deadlineOffset: 1,
    status: "submitted",
    submissionContent:
      "Implementasi endpoint CRUD pengguna beserta validasi input.",
    submissionUrl: "https://github.com/bagas/crud-users",
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "backend-engineering",
    materialIndex: 4,
    materialTitle: "Studi Kasus",
    position: 4,
    deadlineOffset: 8,
    status: "not_started",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "ui-ux-design",
    materialIndex: 1,
    materialTitle: "User Research",
    position: 1,
    deadlineOffset: -4,
    status: "submitted",
    submissionContent: "Laporan wawancara 5 responden beserta persona.",
    submissionUrl: "https://www.figma.com/file/contoh",
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "ui-ux-design",
    materialIndex: 2,
    materialTitle: "Wireframe",
    position: 2,
    deadlineOffset: 0,
    status: "in_progress",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "ui-ux-design",
    materialIndex: 3,
    materialTitle: "Prototyping",
    position: 3,
    deadlineOffset: 12,
    status: "not_started",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "product-management",
    materialIndex: 1,
    materialTitle: "Riset Pengguna",
    position: 1,
    deadlineOffset: -10,
    status: "completed",
    submissionContent: "PRD awal untuk fitur onboarding.",
    submissionUrl: null,
    reviewDecision: "approved",
    reviewNote: "Tujuan dan metrik sudah jelas.",
    reviewerName: "Dinda Maharani",
  },
  {
    cohortId: "product-management",
    materialIndex: 2,
    materialTitle: "Penyusunan PRD",
    position: 2,
    deadlineOffset: -1,
    status: "not_started",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "data-science",
    materialIndex: 1,
    materialTitle: "Eksplorasi Dataset",
    position: 1,
    deadlineOffset: 3,
    status: "in_progress",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
  {
    cohortId: "data-science",
    materialIndex: 2,
    materialTitle: "Visualisasi Data",
    position: 2,
    deadlineOffset: 15,
    status: "not_started",
    submissionContent: null,
    submissionUrl: null,
    reviewDecision: null,
    reviewNote: null,
    reviewerName: null,
  },
];
